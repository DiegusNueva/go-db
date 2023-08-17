package main

import (
	"github.com/DiegusNueva/go-db/pkg/storage"
)

func main() {
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
}
