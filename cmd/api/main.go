package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cthiagoodev/thiagoodev-portfolio/server/internal/common"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	config := common.NewConfig()
	pool, err := common.NewDb(config.DbUrl)

	if err != nil {
		log.Fatal("error on init server: ", err)
		return
	}

	router := NewRouter(pool)

	fmt.Print("Starting server on port " + config.Port + "...")
	serverError := http.ListenAndServe(":"+config.Port, router)

	if serverError != nil {
		log.Fatal("error on init server: ", serverError)
	}
}
