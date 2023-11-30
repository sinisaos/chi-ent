package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/sinisaos/chi-ent/ent"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open new connection
func Open(dsn string) *ent.Client {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	// Cache
	driver := entcache.NewDriver(
		drv,
		entcache.TTL(time.Second),
		entcache.Levels(entcache.NewLRU(128)),
	)
	return ent.NewClient(ent.Driver(driver))
}

func DbConnection() *ent.Client {
	// Create ent client.
	client := Open(Config("DSN"))
	// Run the migrations.
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	return client
}
