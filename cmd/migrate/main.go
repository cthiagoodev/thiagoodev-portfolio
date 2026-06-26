package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	// Driver de source para o migrate, permite ler arquivos de migração do disco.
	// O import em branco (underscore) é necessário para que o driver se registre
	// com a biblioteca principal do migrate através de sua função init().
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// Driver de banco de dados que atua como uma ponte (adapter) entre a interface
	// padrão `database/sql` e a biblioteca de alta performance `pgx`.
	// O import em branco registra o driver "pgx" para ser usado pela função sql.Open.
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

const EnvFileName = ".env"
const EnvDatabaseKey = "DATABASE_URL"
const DriverName = "pgx"
const MigrationsPath = "file://migrations"
const DatabaseDriverName = "postgres"

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)

var ErrNoDatabaseEnvironmentVariable = errors.New("DATABASE_URL environment variable is not set")

func main() {
	godotenv.Load(EnvFileName)

	db, err := getDbInstance()
	if err != nil {
		log.Fatalf("%sCould not get database instance: %v%s", ColorRed, err, ColorReset)
	}

	defer db.Close()

	mgt, err := getMigrationInstance(db)
	if err != nil {
		log.Fatalf("%sCould not get migration instance: %v%s", ColorRed, err, ColorReset)
	}

	defer func() {
		if srcErr, dbErr := mgt.Close(); srcErr != nil || dbErr != nil {
			log.Printf("%sError closing migration instance: src_err=%v, db_err=%v%s", ColorYellow, srcErr, dbErr, ColorReset)
		}
	}()

	if len(os.Args) < 2 {
		log.Fatalf("%sMissing command. Usage: go run ./cmd/migrate/main.go [up|down]%s", ColorRed, ColorReset)
	}

	cmd := os.Args[1]
	switch cmd {
	case "up":
		Up(mgt)
	case "down":
		Down(mgt)
	default:
		log.Fatalf("%sInvalid command: '%s'. Use 'up' or 'down'%s", ColorRed, cmd, ColorReset)
	}
}

func Up(mgt *migrate.Migrate) {
	if err := mgt.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("%sError applying 'up' migration: %v%s", ColorRed, err, ColorReset)
	}
	log.Printf("%s'Up' migration applied successfully!%s", ColorGreen, ColorReset)
}

func Down(mgt *migrate.Migrate) {
	if err := mgt.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("%sError reverting 'down' migration: %v%s", ColorRed, err, ColorReset)
	}
	log.Printf("%s'Down' migration reverted successfully!%s", ColorGreen, ColorReset)
}

func getDbInstance() (*sql.DB, error) {
	url := os.Getenv(EnvDatabaseKey)
	if url == "" {
		return nil, ErrNoDatabaseEnvironmentVariable
	}
	return sql.Open(DriverName, url)
}

func getMigrationInstance(db *sql.DB) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	return migrate.NewWithDatabaseInstance(MigrationsPath, DatabaseDriverName, driver)
}
