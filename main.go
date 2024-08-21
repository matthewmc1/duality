package main

import (
	"log"

	"github.com/mmcgibbon1/duality/data"
)

func main() {
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
}
