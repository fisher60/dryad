package database

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/fisher60/dryad/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var Engine *Queries

//go:embed migrations/*.sql
var embedMigrations embed.FS

func createConnString(config config.DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		config.Host, config.Port, config.User, config.Password, config.DbName,
	)
}

func InitializeDatabse(config config.DatabaseConfig) {
	dbpool, err := pgxpool.New(context.Background(), createConnString(config))
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Print("Migrating database")

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	db := stdlib.OpenDBFromPool(dbpool)
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	Engine = New(dbpool)
}
