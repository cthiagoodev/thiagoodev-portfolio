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

	db, dbError := getDbInstance()

	if dbError != nil {
		log.Fatal("Error on get database instance")
		return
	}

	dvr, dvrError := getDbDriver(db)

	if dvrError != nil {
		log.Fatal("Error on get driver instance")
		return
	}

	mgt, mgtError := getMigrationInstance(dvr)

	if mgtError != nil {
		log.Fatal("Error on get migration instance")
		return
	}

	cmd := args[0]

	switch cmd {
	case "up":
		Up(mgt)
	case "down":
		Down(mgt)
	default:
		log.Fatal("Invalid args")
	}

	defer func(mgt *migrate.Migrate) {
		clsError := db.Close()

		if clsError != nil {
			log.Fatal("Error on close database")
			return
		}
	}(mgt)
}

func Up(mgt *migrate.Migrate) {
	err := mgt.Up()

	if err != nil {
		log.Printf("Error on migrate up")
	}

	log.Println("Migration up done")
}

func Down(mgt *migrate.Migrate) {
	err := mgt.Down()

	if err != nil {
		log.Printf("Error on migrate down")
	}

	log.Println("Migration down done")
}

func getDbInstance() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")
	return sql.Open("pgx", url)
}

func getDbDriver(db *sql.DB) (database.Driver, error) {
	return postgres.WithInstance(db, &postgres.Config{})
}

func getMigrationInstance(dvr database.Driver) (*migrate.Migrate, error) {
	return migrate.NewWithDatabaseInstance("file://migrations", "postgres", dvr)
}
