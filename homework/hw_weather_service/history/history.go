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

var (
	es1 = os.Getenv("ES01")
	es2 = os.Getenv("ES02")
	es3 = os.Getenv("ES03")
)

// Client ...
type Client struct {
	// HTTP client used to make requests.
	*elasticsearch.Client
	IndexName string
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
			fmt.Printf("Attempt: %duration | Sleeping for %s...\n", i, duration)
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

// Save ...
func (c *Client) Save(ind string, r string) (rtn map[string]interface{}, err error) {
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
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status())
	} else {
		// Response into a map.
		if err := json.NewDecoder(res.Body).Decode(&rtn); err != nil {
			log.Printf("error: json.NewDecoder(res.Body): %s", err)
		} else {
			// Response status
			log.Printf(
				"[%s] %s; version=%d",
				res.Status(),
				rtn["result"],
				int(rtn["_version"].(float64)))
		}
	}
	return
}

// SearchForEntries ...
func (c *Client) SearchForEntries(querySearch string) (strRes string, err error) {
	// Compose request body
	var body strings.Builder
	body.WriteString(`
	{
	  "query": {
		"multi_match" : {
		  "query":    "` + querySearch + `",
		  "fields": [ "_index", "title" , "traceID" , "timestamp" , "body" ]
		}
	  }
	}`)

	var res *esapi.Response
	res, err = c.Search(
		c.Search.WithIndex(c.IndexName),
		c.Search.WithBody(strings.NewReader(body.String())),
		c.Search.WithTrackTotalHits(true),
		c.Search.WithSize(10),
		c.Search.WithPretty(),
		c.Search.WithFilterPath("took", "hits.hits"),
	)

	strRes = res.String()
	log.Println("\x1b[1mResponse:\x1b[0m", strRes)
	if len(strRes) <= len("[200 OK] ") {
		log.Printf("Response body is empty")
	}

	if err != nil {
		log.Printf("Error:   %strRes", err)
		return
	}

	return
}

// // Historiador ...
// type Historiador interface {
// 	Store(record *Record) (rtn string, err error)
// 	ShowStored()
// }

// // Record ...
// type Record struct {
// 	TraceID   string
// 	LogString map[string]interface{}
// }
//
// // NewRecord ...
// func NewRecord(traceID string, record map[string]interface{}) *Record {
// 	return &Record{TraceID: traceID, LogString: record}
// }
