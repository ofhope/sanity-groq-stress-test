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

type QueryResult struct {
	Id   string
	Time string
}

func (q *QueryResult) String() string {
	return fmt.Sprintf("%s, %s \n", q.Id, q.Time)
}

func NewClient() Client {
	projectId := os.Getenv("SANITY_STUDIO_PROJECT_ID")
	dataset := os.Getenv("SANITY_STUDIO_API_DATASET")

	c, err := sanity.New(projectId,
		sanity.WithCallbacks(sanity.Callbacks{
			OnQueryResult: func(result *sanity.QueryResult) {
				log.Printf("Sanity queried in %d ms!", result.Time)
			},
		}),
		sanity.WithDataset(dataset))

	if err != nil {
		log.Fatal(err)
	}

	return Client{
		instance: c,
	}
}

func (c *Client) RunQuery(rawQuery string, ch chan<- QueryResult) {
	query := c.instance.Query(rawQuery)

	result, err := query.Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	ch <- QueryResult{
		Id:   "",
		Time: result.Time.String(),
	}
}
