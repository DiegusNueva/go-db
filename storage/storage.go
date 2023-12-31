package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {

	once.Do(func() {
		//here, one time
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:Ereba555@@172.17.0.2:5432/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("Conectado a postgres")

	})

}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}
