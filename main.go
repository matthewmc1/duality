package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mmcgibbon1/duality/data"
)

func main() {
	ctx := context.Background()

	dbclient := data.DbClient{
		Dbtype: "duckdb",
		Dbpath: "/Users/mmcgibbon/duality/data/duckdb/evergreen.db",
	}
	db, err := data.Client(dbclient)
	if err != nil {
		log.Println("error creating client", err)
	}
	err = data.RunMigrations(db, "/Users/mmcgibbon/duality/data/migrations")
	if err != nil {
		log.Println("error running migrations", err)
	}

	reqctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, err := uuid.NewV7()
	if err != nil {
		log.Println("error creating uuid", err)
	}

	eg := data.Evergreen{
		Id:          id,
		Title:       "Testing async flows",
		Label:       "test",
		CreatedDate: time.Now().Format(time.RFC3339),
		Details:     "Sync",
	}

	id2, err := uuid.NewV7()
	if err != nil {
		log.Println("error creating uuid", err)
	}

	eg2 := data.Evergreen{
		Id:          id2,
		Title:       "Testing async flows new",
		Label:       "test",
		CreatedDate: time.Now().Format(time.RFC3339),
		Details:     "Sync",
	}

	dbclient.Create(reqctx, db, eg)
	dbclient.DeleteByTitle(reqctx, db, eg2.Title)
}
