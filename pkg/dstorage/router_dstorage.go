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
	r.Path("/download/{name}").Methods("GET").HandlerFunc(getFile(ds))
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
		file, header, err := r.FormFile("fileupload")
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

		if err := ds.UploadFile(buf.Bytes(), header.Filename); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error uploading file: %v", err)
		}
	}
}

func getFile(ds *DStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		file, err := ds.GetFile(vars["name"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error uploading file: %v", err)
		}
		// не помню, как вернуть, гуглить лень
		fmt.Println(file)
	}
}
