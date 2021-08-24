package history

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

// Record ...
type Record struct {
	TraceID   string
	LogString map[string]interface{}
}
type Client struct {
	// HTTP client used to make requests.
	*elasticsearch.Client
}

func NewClient() *Client {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	return &Client{Client: es}
}

func (c *Client) WriteToIndex(indexName string, text string) map[string]interface{} {
	// Build the request body.
	var b strings.Builder
	b.WriteString(`{"title" : "`)
	b.WriteString(text)
	b.WriteString(`"}`)

	// Set up the request object.
	req := esapi.IndexRequest{
		Index: indexName,
		// DocumentID: strconv.Itoa(i + 1),
		Body:    strings.NewReader(b.String()),
		Refresh: "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), c)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	var r map[string]interface{}
	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status())
	} else {
		// Deserialize the response into a map.
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	return r
}

// // Historiador ...
// type Historiador interface {
// 	Store(record *Record) (rtn string, err error)
// 	ShowStored()
// }
