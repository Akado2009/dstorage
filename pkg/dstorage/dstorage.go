package dstorage

import (
	"sync"

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

func (ds *DStorage) UploadFile(file []byte, name string) error {
	chunks := splitFile(file, len(ds.Servers))

	errs := make(chan error, len(chunks))
	var wg sync.WaitGroup
	for i, v := range chunks {
		wg.Add(1)
		go ds.Servers[i].PutPart(v, name, &errs)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		if err != nil {
			return err
		}
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
	file := mergeChunks(&chunks)
	return file, nil
}
