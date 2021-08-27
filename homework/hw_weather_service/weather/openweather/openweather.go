package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	. "github.com/rodkevich/go-course/homework/hw_weather_service/weather/types"
)

var (
	request      *http.Request
	response     *http.Response
	responseBody []byte
)

// Client ...
type Client struct {
	*http.Client
	BaseURL   string
	UserAgent string
	apiKey    string
	units     string
}

// NewOpenWeatherClient ...
func NewOpenWeatherClient(baseURL string, userAgent string, apiKey string, units string) *Client {
	return &Client{
		Client:    &http.Client{},
		BaseURL:   baseURL,
		UserAgent: userAgent,
		apiKey:    apiKey,
		units:     units,
	}
}

// GetByCityName ...
func (cl *Client) GetByCityName(name string) (weather string, err error) {
	method := "GET"
	url := fmt.Sprintf(
		"%s/data/2.5/weather",
		cl.BaseURL,
	)

	request, err = http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	que := request.URL.Query()
	que.Add("q", name)
	que.Add("appid", cl.apiKey)
	que.Add("units", cl.units)
	request.URL.RawQuery = que.Encode()

	response, err = cl.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	responseBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var in OpenWeatherResponse
	err = json.Unmarshal(responseBody, &in)
	if err != nil {
		return
	}

	tr := &TemperatureResponse{
		CityID:        in.Id,
		City:          in.Name,
		TimeRequested: time.Now().UTC().Format("2006-01-02T15:04:05.999Z"),
		Temperature:   in.Main.Temp,
	}

	out, err := json.Marshal(tr)
	if err != nil {
		panic(err)
	}
	weather = string(out)
	return
}
