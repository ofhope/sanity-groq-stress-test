package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	envFile := flag.String("e", ".env", "optionally specify an env file with `=example.env`")

	err := godotenv.Load(*envFile)

	if err != nil {
		log.Fatalf("No env file found")
	}
}
