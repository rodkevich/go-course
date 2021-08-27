package main

import (
	"log"
	"os"

	"github.com/rodkevich/go-course/homework/hw_weather_service/weather"
)

var (
	weatherPort = os.Getenv("WEATHERPORT")
)

func main() {
	r := weather.SetupService()
	err := r.Run(":" + weatherPort)
	if err != nil {
		log.Fatal(err)
	}
}
