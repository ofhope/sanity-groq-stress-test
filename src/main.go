package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ofhope/sanity-groq-stress-test/src/lib"
)

func processFile(path string, client lib.Client) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Can't read file path %s", path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		go client.RunQuery(scanner.Text())
	}
}

func main() {
	envFile := *flag.String("e", ".env", "optionally specify an env file with `example.env`")
	flag.Parse()

	inputFile := flag.Arg(0)
	if inputFile == "" {
		fmt.Println("You must pass a text file as an argument.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("No env file found. Please include one with SANITY config.")
	}

	client := lib.NewClient()
	processFile(inputFile, client)
	for {
	}
}
