package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rodkevich/go-course/homework/hw_weather_service/weather/types"
)

var (
	req          *http.Request
	response     *http.Response
	responseBody []byte
	res          types.WeatherApiResponse
)

// Client ...
type Client struct {
	// HTTP client used to make requests.
	*http.Client
	// BaseURL   *url.URL
	BaseURL   string
	UserAgent string

	apiKey string
	units  string
}

// NewOpenWeatherClient ...
func NewOpenWeatherClient(baseURL string, userAgent string, apiKey string, units string) *Client {
	return &Client{Client: &http.Client{}, BaseURL: baseURL, UserAgent: userAgent, apiKey: apiKey, units: units}
}

// GetByCityName ...
func (cl *Client) GetByCityName(cityName string) (rtn types.TemperatureResponse, err error) {
	method := "GET"
	url := fmt.Sprintf("%s/data/2.5/weather?que=%s", cl.BaseURL, cityName)
	req, err = http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	que := req.URL.Query()
	que.Add("que", cityName)
	que.Add("appid", cl.apiKey)
	que.Add("units", cl.units)
	req.URL.RawQuery = que.Encode()

	response, err = cl.Do(req)
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
	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		return
	}
	rtn.CityID = res.Id
	rtn.City = res.Name
	rtn.TimeRequested = time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	rtn.Temperature = res.Main.Temp
	return
}
