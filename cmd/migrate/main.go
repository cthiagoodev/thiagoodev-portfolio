// O pacote main para a ferramenta de linha de comando (CLI) de migração de banco de dados.
//
// Este programa é um executável isolado, responsável exclusivamente por aplicar
// ou reverter migrações de esquema do banco de dados. Ele é projetado para ser
// executado manualmente durante o desenvolvimento ou como um passo em um pipeline de CI/CD
// antes de implantar a aplicação principal.
//
// Uso:
//
//	go run ./cmd/migrate/main.go [comando]
//
// Comandos disponíveis:
//
//	up:   Aplica todas as migrações pendentes.
//	down: Reverte a última migração aplicada.
package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
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

// main é o ponto de entrada da aplicação.
// Ele orquestra a configuração, a análise dos argumentos de linha de comando
// e a execução da ação de migração solicitada (up ou down).
func main() {
	// Carrega as variáveis de ambiente de um arquivo .env para fácil configuração
	// em ambiente de desenvolvimento.
	godotenv.Load(".env")

	// --- Cadeia de Instanciação ---
	// 1. Obtém uma instância de conexão com o banco de dados.
	db, dbError := getDbInstance()
	if dbError != nil {
		log.Fatal("Erro ao obter a instância do banco de dados: ", dbError)
	}

	// 2. Cria um driver de banco de dados para o migrate a partir da conexão existente.
	dvr, dvrError := getDbDriver(db)
	if dvrError != nil {
		log.Fatal("Erro ao obter o driver de banco de dados para o migrate: ", dvrError)
	}

	// 3. Cria a instância principal do migrate.
	mgt, mgtError := getMigrationInstance(dvr)
	if mgtError != nil {
		log.Fatal("Erro ao obter a instância principal do migrate: ", mgtError)
	}

	// Garante que a conexão com o banco de dados e os recursos do migrate
	// sejam liberados quando a função main terminar, usando defer.
	defer func(mgt *migrate.Migrate) {
		if clsMgtError, dbMgtError := mgt.Close(); clsMgtError != nil || dbMgtError != nil {
			log.Fatal("Erro ao fechar os recursos do migrate: ", clsMgtError, dbMgtError)
		}
	}(mgt)

	// --- Lógica de Execução ---
	// Pula o primeiro argumento, que é o nome do próprio programa.
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Comando de migração ausente. Use 'up' ou 'down'.")
	}

	cmd := args[0]
	switch cmd {
	case "up":
		Up(mgt)
	case "down":
		Down(mgt)
	default:
		log.Fatalf("Comando inválido: '%s'. Use 'up' ou 'down'.", cmd)
	}
}

// Up aplica todas as migrações pendentes.
// Ele ignora o erro `migrate.ErrNoChange`, que ocorre quando não há
// novas migrações a serem aplicadas, tratando-o como um sucesso.
func Up(mgt *migrate.Migrate) {
	if err := mgt.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Erro ao aplicar a migração 'up': %v", err)
	}
	log.Println("Migração 'up' concluída com sucesso.")
}

// Down reverte a última migração aplicada.
// Assim como a função Up, ele trata `migrate.ErrNoChange` como um caso de sucesso.
func Down(mgt *migrate.Migrate) {
	if err := mgt.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Erro ao reverter a migração 'down': %v", err)
	}
	log.Println("Migração 'down' concluída com sucesso.")
}

// getDbInstance abre e retorna uma conexão com o banco de dados usando a interface
// padrão `database/sql`.
// Ele usa o driver "pgx" (registrado via blank import) para se conectar ao
// PostgreSQL através da URL fornecida na variável de ambiente DATABASE_URL.
func getDbInstance() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")
	return sql.Open("pgx", url)
}

// getDbDriver cria e retorna um driver de banco de dados específico para o
// `golang-migrate` a partir de uma conexão `*sql.DB` existente.
func getDbDriver(db *sql.DB) (database.Driver, error) {
	return postgres.WithInstance(db, &postgres.Config{})
}

// getMigrationInstance cria e retorna a instância principal do `migrate`.
// Esta instância é configurada para:
//   - Ler arquivos de migração do diretório "migrations" (source: "file://...").
//   - Conectar-se a um banco de dados do tipo "postgres" (database name).
//   - Usar o driver de banco de dados fornecido.
func getMigrationInstance(dvr database.Driver) (*migrate.Migrate, error) {
	return migrate.NewWithDatabaseInstance("file://migrations", "postgres", dvr)
}
