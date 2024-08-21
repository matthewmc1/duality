package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/marcboeker/go-duckdb"
)

type DbClient struct {
	Dbtype   string
	Dbpath   string //used for duckdb
	Dbname   string
	Username string //optional
	Password string //optional
	Port     string //optional
	Protocol string //optional:SSL,PLAINTEXT
	Sslmode  bool   //optional
}

func RunMigrations(db *sqlx.DB, migrationsDir string) error {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("error reading migrations directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			migrationFilePath := filepath.Join(migrationsDir, file.Name())
			migrationSQL, err := os.ReadFile(migrationFilePath)
			if err != nil {
				return fmt.Errorf("error reading migration file %s: %v", file.Name(), err)
			}

			_, err = db.Exec(string(migrationSQL))
			if err != nil {
				return fmt.Errorf("error executing migration %s: %v", file.Name(), err)
			}

			fmt.Printf("Successfully executed migration: %s\n", file.Name())
		}
	}

	return nil
}

func Client(client DbClient) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	if client.Dbtype == "duckdb" {
		db, err = sqlx.Connect(client.Dbtype, client.Dbpath)
	} else {
		db, err = sqlx.Connect(client.Dbtype, "")
	}
	if err != nil {
		log.Fatalln(err)
		return nil, errors.New("error creating db")
	}

	return db, nil
}

type Database interface {
	Create() error
	Delete(id string) error
}

func (db *DbClient) Create(ctx context.Context, client *sqlx.DB, event Evergreen) error {
	tx := client.MustBeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})

	res, err := tx.NamedExec("INSERT INTO evergreen VALUES (:id, :title, :labels, :created_date, :details)", &event)
	if err != nil {
		log.Fatalln(res, err)
		return err
	}

	tx.Commit()
	return nil
}
