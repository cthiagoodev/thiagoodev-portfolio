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
	godotenv.Load(".env")

	config := common.NewConfig()
	pool, err := common.NewDb(config.DbUrl)

	defer pool.Close()

	if err != nil {
		log.Fatal("error on init server: ", err)
		return
	}

	tm, tmErr := templates.NewTemplateManager()

	if tmErr != nil {
		log.Fatal("error on init server: ", tmErr)
		return
	}

	router := NewRouter(pool, tm)

	fmt.Print("Starting server on port " + config.Port + "...")
	serverError := http.ListenAndServe(":"+config.Port, router)

	if serverError != nil {
		log.Fatal("error on init server: ", serverError)
	}
}
