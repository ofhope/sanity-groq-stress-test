package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	sanity "github.com/sanity-io/client-go"
)

func main() {

	envFile := flag.String("e", ".env", "optionally specify an env file with `example.env`")

	err := godotenv.Load(*envFile)
	if err != nil {
		log.Fatalf("No env file found")
	}

	projectId := os.Getenv("SANITY_STUDIO_PROJECT_ID")
	dataset := os.Getenv("SANITY_STUDIO_API_DATASET")
	token := os.Getenv("SANITY_STUDIO_TOKEN")

	client, err := sanity.New(projectId,
		sanity.WithCallbacks(sanity.Callbacks{
			OnQueryResult: func(result *sanity.QueryResult) {
				log.Printf("Sanity queried in %d ms!", result.Time)
			},
		}),
		sanity.WithToken(token),
		sanity.WithDataset(dataset))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to project:%s dataset:%s\n", projectId, dataset)

	rawQuery := `
		*[]{
			_id,
			body[_type match $type]
		}
	`
	query := client.Query(rawQuery)

	result, err := query.Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Time %d \n", result.Time)
	bytes, err := result.Result.MarshalJSON()
	fmt.Printf("Result %s \n", string(bytes))
}
