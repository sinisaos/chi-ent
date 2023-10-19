package database

import (
	"context"
	"log"

	"github.com/sinisaos/chi-ent/ent"

	_ "github.com/mattn/go-sqlite3"
)

func DbConnection() *ent.Client {
	// Create ent client.
	client, err := ent.Open("sqlite3", Config("DSN"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database is connected.")
	// Run the migrations.
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	return client
}
