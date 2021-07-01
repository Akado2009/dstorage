package dstorage

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

type server struct {
	Path string
}

func (s *server) CheckAndSetup() error {
	if _, err := os.Stat(s.Path); os.IsNotExist(err) {
		// create
		return os.MkdirAll(s.Path, os.ModePerm)
	} else {
		return nil
	}
}

func (s *server) PutPart(part []byte, name string) error {
	return ioutil.WriteFile(fmt.Sprintf("%s", name), part, fs.ModePerm)
}
