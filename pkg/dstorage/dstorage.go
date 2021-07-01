package dstorage

import db "github.com/Akado2009/dstorage/pkg/mongo"

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

func (ds *DStorage) UploadFile(file []byte) error {
	return nil
}

func (ds *DStorage) GetFile() ([]byte, error) {
	return nil, nil
}
