package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Invalid args")
		return
	}

	cmd := args[0]

	switch cmd {
	case "up":
		Up()
	case "down":
		Down()
	default:
		log.Fatal("Invalid args")
	}
}

func Up() {

}

func Down() {

}

func getDbInstance() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	conn, err := sql.Open("pgx", url)

	if err != nil {
		log.Fatal("Error on connection database")
		return nil
	}

	return conn
}

func getDbDriver(db *sql.DB) *database.Driver {
	dvr, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatal("Error on connection database")
		return nil
	}

	return &dvr
}

func getMigrationInstance(dvr database.Driver) *migrate.Migrate {
	mgt, err := migrate.NewWithDatabaseInstance("file://migrations", "pgx", dvr)

	if err != nil {
		log.Fatal("Error on connection database")
		return nil
	}

	return mgt
}
