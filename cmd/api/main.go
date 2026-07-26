package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common"
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common/templates"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error on init server: ", err)
	}

	config := common.NewConfig()
	pool, err := common.NewDb(config.DbUrl)

	if err != nil {
		log.Fatal("error on init server: ", err)
	}

	defer pool.Close()

	tm, tmErr := templates.NewTemplateManager()

	if tmErr != nil {
		log.Fatal("error on init server: ", tmErr)
	}

	router := NewRouter(pool, tm)

	fmt.Print("Starting server on port " + config.Port + "...")
	serverError := http.ListenAndServe(":"+config.Port, router)

	if serverError != nil {
		log.Fatal("error on init server: ", serverError)
	}
}
