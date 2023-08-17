package main

import (
	"github.com/DiegusNueva/go-db/storage"
)

func main() {
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
}
