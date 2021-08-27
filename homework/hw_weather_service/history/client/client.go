package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/estransport"
)

var (
	es1 = os.Getenv("ES01")
	es2 = os.Getenv("ES02")
	es3 = os.Getenv("ES03")
)

// Client ...
type Client struct {
	// HTTP client used to make requests.
	IndexName string
	*elasticsearch.Client
}

// NewEsClient ...
func NewEsClient(indexName string) *Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			es1,
			es2,
			es3,
			"http://localhost:9200",
			"http://localhost:9300",
		},
		RetryOnStatus: []int{429, 502, 503, 504},
		RetryBackoff: func(i int) time.Duration {
			duration := time.Duration(math.Exp2(float64(i))) * 3 * time.Second
			fmt.Printf("Attempt: %v duration | Sleeping for %s...\n", i, duration)
			return duration
		},
		Logger: &estransport.JSONLogger{Output: os.Stdout},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
		return nil
	}
	return &Client{Client: es, IndexName: indexName}
}

// SaveWithIndex ...
func (c *Client) SaveWithIndex(ind string, r string) (rtn *map[string]interface{}, err error) {
	// Create request object
	req := esapi.IndexRequest{
		Index:   ind,
		Body:    strings.NewReader(r),
		Refresh: "true",
		Pretty:  true,
	}

	// Perform the request
	res, err := req.Do(context.Background(), c)
	if err != nil {
		log.Printf("error: getting response for SaveWithIndex: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("error: history.client: error indexing document ID=%s", res.Status())
	} else {
		// Response into a map to be returned
		if err := json.NewDecoder(res.Body).Decode(&rtn); err != nil {
			log.Printf("error: history.client: json.NewDecoder(res.Body): %s", err)
		} else {
			// Log response status
			log.Printf(
				"[%s] %s",
				res.Status(),
				ind+": history recodrd created",
			)
		}
	}
	return rtn, nil
}

// SearchForEntries ...
func (c *Client) SearchForEntries(querySearch string) (entries string, err error) {
	// Compose request body
	var body strings.Builder
	body.WriteString(`
	{
	  "query": {
		"multi_match" : {
		  "query":    "` + querySearch + `",
		  "fields": [ "_index", "title" , "traceID", "body" ]
		}
	  }
	}`)

	var res *esapi.Response
	res, err = c.Search(
		c.Search.WithBody(strings.NewReader(body.String())),
		c.Search.WithTrackTotalHits(true),
		c.Search.WithSize(20),
		c.Search.WithPretty(),
		c.Search.WithFilterPath("took", "hits.hits"),
	)
	if err != nil {
		log.Printf("error: %entries", err)
	}

	entries = res.String()
	if len(entries) <= len("[200 OK] ") {
		log.Printf("Response body is empty")
	}

	return entries, nil
}
