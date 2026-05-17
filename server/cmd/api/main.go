package main

import (
	"log"
	"net/http"

	"github.com/cthiagoodev/thiagoodev-portfolio/server/internal/common"
)

func main() {
	config := common.NewConfig()
	pool, err := common.NewDb(config.DbUrl)

	if err != nil {
		log.Fatal("error on init server: ", err)
		return
	}

	router := NewRouter(pool)

	http.ListenAndServe(":"+config.Port, router)
}
