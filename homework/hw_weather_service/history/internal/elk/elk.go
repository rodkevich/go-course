package elk

import (
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"

	"github.com/rodkevich/go-course/homework/hw_weather_service/history"
)


type historyService struct {
	es *history.Client
}

// Store ...
func (h historyService) Store(record *history.Record) (rtn string, err error) {
	r, _ := h.es.Index(
		"weather_service_logging",
		strings.NewReader(`{"title" : "weather_service_logging"}`),
		h.es.Index.WithRefresh("true"),
		h.es.Index.WithPretty(),
		h.es.Index.WithFilterPath("result", "_id"),
	)
	return r.String(), nil
}

// ShowStored ...
func (h historyService) ShowStored() {
	panic("implement me")
}

// NewHistoryService ...
func NewHistoryService() (history.Historiador, error) {
	log.SetFlags(0)
	var es *elasticsearch.Client
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &estransport.JSONLogger{Output: os.Stdout},
	})
	return &historyService{es: es}, nil
}

func main() {
	log.SetFlags(0)

	var es *elasticsearch.Client

	// ==============================================================================================
	//
	// "JSONLogger" writes the information as JSON and is suitable for production logging.
	//
	es, _ = elasticsearch.NewClient(elasticsearch.Config{
		Logger: &estransport.JSONLogger{Output: os.Stdout},
	})
	run(es, "JSON")
}

// ------------------------------------------------------------------------------------------------

func run(es *elasticsearch.Client, name string) {
	// es.Delete("test", "1")
	// es.Exists("test", "1")
	// es.Search(es.Search.WithQuery("{FAIL"))
	res, err := es.Search(
		es.Search.WithIndex("test"),
		es.Search.WithBody(strings.NewReader(`{"query" : {"match" : { "title" : "weather_service_logging" } } }`)),
		es.Search.WithSize(1),
		es.Search.WithPretty(),
		es.Search.WithFilterPath("took", "hits.hits"),
	)

	s := res.String()
	// log.Println("\x1b[1mResponse:\x1b[0m", s)
	if len(s) <= len("[200 OK] ") {
		log.Fatal("Response body is empty")
	}

	if err != nil {
		log.Fatalf("Error:   %s", err)
	}

	log.Print("\n")
}
