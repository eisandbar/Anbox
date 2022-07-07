package database

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"

	typing "eisandbar/anbox/typing"
)

var (
	HOST        = os.Getenv("POSTGRES_HOST")
	USER        = os.Getenv("POSTGRES_USER")
	DB_NAME     = os.Getenv("POSTGRES_DB")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
)

var connStr = fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable",
	HOST,
	USER,
	DB_NAME,
	DB_PASSWORD,
)

func (repo *Repository) InitDB() {
	// Migrating schema
	repo.db.AutoMigrate(&typing.Game{})
	repo.db.AutoMigrate(&typing.User{})
	repo.db.AutoMigrate(&typing.Link{})

	fmt.Println("Postgres DB initialized")
}
