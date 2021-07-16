package lib

import (
	"context"
	"fmt"
	"log"
	"os"

	sanity "github.com/sanity-io/client-go"
)

type Client struct {
	instance *sanity.Client
}

func NewClient() Client {
	projectId := os.Getenv("SANITY_STUDIO_PROJECT_ID")
	dataset := os.Getenv("SANITY_STUDIO_API_DATASET")
	token := os.Getenv("SANITY_STUDIO_TOKEN")

	c, err := sanity.New(projectId,
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

	return Client{
		instance: c,
	}
}

func (c *Client) RunQuery(rawQuery string, ch chan<- string) {
	query := c.instance.Query(rawQuery)

	result, err := query.Do(context.Background())
	if err != nil {
		ch <- fmt.Sprintf("Error while running query %v", err)
		log.Fatal(err)
	}

	ch <- result.Time.String()
}
