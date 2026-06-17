package main

import (
	"log"
	"os"

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
