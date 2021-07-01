package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Akado2009/dstorage/pkg/dstorage"
	"github.com/Akado2009/dstorage/pkg/models"
	db "github.com/Akado2009/dstorage/pkg/mongo"
)

func main() {
	config, err := models.LoadConfig("")
	if err != nil {
		log.Fatalf("Failed to load config file: %v\n", err)
	}
	log.Println("Config loaded successfully...")
	log.Printf("%#v\n", config)
	// create db to storage
	// storageDB :=
	storageDB, err := db.NewDatabase(&config.DB)
	if err != nil {
		log.Fatalf("Failed to setup a db connection: %v\n", err)
	}
	log.Println("DB connectioon established...")

	ss := dstorage.NewDStorage(storageDB)
	log.Println("DStorage created...")

	wss := dstorage.NewRouterDStorage(ss)
	log.Println("Router DStorage created...")

	srv := &http.Server{
		Handler: wss.Router,
		Addr:    fmt.Sprintf(":%d", config.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
