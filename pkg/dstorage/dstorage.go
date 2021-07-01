package dstorage

import (
	"sync"

	"github.com/Akado2009/dstorage/pkg/models"
	db "github.com/Akado2009/dstorage/pkg/mongo"
)

type DStorage struct {
	Servers []*server
	DB      *db.DB
}

func NewDStorage(db *db.DB) *DStorage {
	return &DStorage{
		Servers: make([]*server, 0),
		DB:      db,
	}
}

func (ds *DStorage) AddServer(s *server) error {
	if err := s.CheckAndSetup(); err != nil {
		return err
	}
	ds.Servers = append(ds.Servers, s)
	return nil
}

// type Info struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty"`
// 	Name      string             `bson:"name,omitempty"`
// 	Servers   []string           `bson:"servers,omitempty"`
// 	CheckSums []string           `bson:"checkSums,omitempty"`
// }
func (ds *DStorage) UploadFile(file []byte, name string) error {
	chunks := splitFile(file, len(ds.Servers))

	errs := make(chan error, len(chunks))
	var wg sync.WaitGroup

	hashes := make([]string, len(chunks), len(chunks))
	for i, v := range chunks {
		wg.Add(1)
		hashes = append(hashes, calculateHash(v))
		go ds.Servers[i].PutPart(v, name, &errs)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			return err
		}
	}
	servers := make([]string, len(ds.Servers), len(ds.Servers))
	for _, v := range ds.Servers {
		servers = append(servers, v.Path)
	}
	info := &models.Info{
		Name:      name,
		Servers:   servers,
		CheckSums: hashes,
	}

	if err := ds.DB.InsertInfo(info); err != nil {
		return err
	}
	return nil
}

func (ds *DStorage) GetFile(name string) ([]byte, error) {

	chunks := make(map[int][]byte, 0)
	var wg sync.WaitGroup
	errs := make(chan error, len(chunks))
	for i, v := range ds.Servers {
		wg.Add(1)
		go v.GetPart(name, i, chunks, &errs)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			return nil, err
		}
	}
	info, err := ds.DB.GetInfo(name)
	if err != nil {
		return nil, err
	}
	check := checkChunks(chunks, info.CheckSums)
	if !check {
		return nil, err
	}
	file := mergeChunks(&chunks)
	return file, nil
}
