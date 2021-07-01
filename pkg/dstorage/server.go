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

func (s *server) PutPart(part []byte, name string, errs *chan error) {
	*errs <- ioutil.WriteFile(fmt.Sprintf("%s", name), part, fs.ModePerm)
}

func (s *server) GetPart(name string, index int, chunks map[int][]byte, errs *chan error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		*errs <- err
		return
	}
	chunks[index] = b
}
