package history

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

// Record ...
type Record struct {
	TraceID   string
	LogString map[string]interface{}
}

// NewRecord ...
func NewRecord(traceID string, record map[string]interface{}) *Record {
	return &Record{TraceID: traceID, LogString: record}
}

type Client struct {
	// HTTP client used to make requests.
	*elasticsearch.Client
	IndexName string
}

// NewEsClient ...
func NewEsClient(indexName string) *Client {
	cfg := elasticsearch.Config{
		RetryOnStatus: []int{429, 502, 503, 504},
		RetryBackoff: func(i int) time.Duration {
			duration := time.Duration(math.Exp2(float64(i))) * time.Second
			fmt.Printf("Attempt: %duration | Sleeping for %s...\n", i, duration)
			return duration
		},
		Logger: &estransport.JSONLogger{Output: os.Stdout},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return &Client{Client: es, IndexName: indexName}
}

// Save ...
func (c *Client) Save(title string) (rtn map[string]interface{}, err error) {
	// Build the request body.
	var b strings.Builder
	b.WriteString(`{"title" : "`)
	b.WriteString(title)
	b.WriteString(`"}`)

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   c.IndexName,
		Body:    strings.NewReader(b.String()),
		Refresh: "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), c)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status())
	} else {
		// Deserialize the response into a map.
		if err := json.NewDecoder(res.Body).Decode(&rtn); err != nil {
			log.Printf("error: json.NewDecoder(res.Body): %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), rtn["result"], int(rtn["_version"].(float64)))
		}
	}
	return
}

// SearchForEntries ...
func (c *Client) SearchForEntries(querySearch string) (strRes string, err error) {
	// es.Delete("test", "1")
	// es.Exists("test", "1")
	// es.Search(es.Search.WithQuery("{FAIL"))
	// Build the request body.
	var b strings.Builder
	b.WriteString(`
	{
	  "query": {
		"multi_match" : {
		  "query":    "` + querySearch + `",
		  "fields": [ "_index", "title" ]
		}
	  }
	}`)
	log.Println(b.String())
	var res *esapi.Response
	res, err = c.Search(
		c.Search.WithIndex(c.IndexName),
		c.Search.WithBody(strings.NewReader(b.String())),
		c.Search.WithTrackTotalHits(true),
		c.Search.WithSize(10),
		c.Search.WithPretty(),
		c.Search.WithFilterPath("took", "hits.hits"),
	)
	log.Println(res)

	strRes = res.String()
	log.Println("\x1b[1mResponse:\x1b[0m", strRes)
	if len(strRes) <= len("[200 OK] ") {
		log.Printf("Response body is empty")
	}
	if err != nil {
		log.Printf("Error:   %strRes", err)
	}
	return
}

// // Historiador ...
// type Historiador interface {
// 	Store(record *Record) (rtn string, err error)
// 	ShowStored()
// }
