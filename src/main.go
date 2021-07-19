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
		log.Fatalf("Failed opening input file: %s", err)
	}
	defer file.Close()

	if godotenv.Load(envFile) != nil {
		log.Fatalf("No env file found. Please include one with SANITY config.")
	}

	ch := make(chan lib.QueryResult)
	client := lib.NewClient()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		go client.RunQuery(scanner.Text(), ch)
	}

	outputFile, err := os.Create("output.csv")
	if err != nil {
		log.Fatalf("Failed createing output file: %s", err)
	}
	defer outputFile.Close()

	fmt.Println("reading from channel")
	outputFile.WriteString("ID, Time\n")

	for result := range ch {
		_, _ = outputFile.WriteString(result.String())
	}
	outputFile.Sync()
}
