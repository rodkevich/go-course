package main

import (
	"log"
	"os"

	"github.com/rodkevich/go-course/homework/hw_weather_service/history"
)

var (
	historyPort = os.Getenv("HISTORYPORT")
)

func main() {
	r := history.SetupService()
	err := r.Run(":" + historyPort)
	if err != nil {
		log.Fatal(err)
	}
}

/*
For NO auth just remove headers. You'll get 401 status code error

GET With auth:

curl --location --request GET 'http://localhost:9091/logs/this%20text%20will%20be%20logged' \
--header 'Authorization: Basic Z29waGVyOmhpc3RvcnlTZXJ2aWNl'

POST With auth:

curl --location --request POST 'http://localhost:9091/logs/create' \
--header 'Authorization: Basic Z29waGVyOmhpc3RvcnlTZXJ2aWNl' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "this text will be logged"
}'

*/
