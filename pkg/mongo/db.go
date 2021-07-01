package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Akado2009/dstorage/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB is a wrapper on top of mongodb
type DB struct {
	client     *mongo.Client
	collection *mongo.Collection
	cfg        *models.DBConfig
}

// NewDatabase returns a DB instance
func NewDatabase(cfg *models.DBConfig) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	URI := fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}

	database := client.Database(cfg.Database)
	collection := database.Collection(cfg.Collection)

	return &DB{
		collection: collection,
		client:     client,
		cfg:        cfg,
	}, nil
}

func (d *DB) InsertInfo(info *models.Info) error {
	ctx := context.Background()
	id, err := d.collection.InsertOne(ctx, info)
	log.Printf("Inserted info for a file: %v with ID: %v\n", info.Name, id)
	return err
}

func (d *DB) GetInfo(name string) (*models.Info, error) {
	var in models.Info
	ctx := context.Background()
	err := d.collection.FindOne(ctx, &bson.M{"name": name}).Decode(&in)
	if err != nil {
		return nil, err
	}
	return &in, err
}

func (d *DB) Close() {
	ctx := context.Background()
	if err := d.client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
