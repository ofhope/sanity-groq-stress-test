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

func main() {
	envFile := *flag.String("e", ".env", "optionally specify an env file with `example.env`")
	flag.Parse()

	inputFile := flag.Arg(0)
	if inputFile == "" {
		fmt.Println("You must pass a text file as an argument.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Can't read file path %s", inputFile)
	}
	defer file.Close()

	if godotenv.Load(envFile) != nil {
		log.Fatalf("No env file found. Please include one with SANITY config.")
	}

	ch := make(chan string)
	client := lib.NewClient()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		go client.RunQuery(scanner.Text(), ch)
	}
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(<-ch)
	}
	for {
	}
}
