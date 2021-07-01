package dstorage

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RouterDStorage struct {
	Router *mux.Router
	ds     *DStorage
}

func NewRouterDStorage(ds *DStorage) *RouterDStorage {
	r := mux.NewRouter()
	r.Path("/upload").Methods("PUT").HandlerFunc(uploadFile(ds))
	return &RouterDStorage{
		Router: r,
		ds:     ds,
	}
}

func uploadFile(ds *DStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := r.ParseMultipartForm(5 * 1024 * 1024)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error uploading file: %v", err)
		}
		file, _, err := r.FormFile("fileupload")
		defer file.Close()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error uploading file: %v", err)
		}
		defer file.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error uploading file: %v", err)
		}
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error uploading file: %v", err)
		}

		if err := ds.UploadFile(buf.Bytes()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error uploading file: %v", err)
		}
	}
}

func uploadFIle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 * 1024 * 1024)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.FormValue("fileupload"))
}
