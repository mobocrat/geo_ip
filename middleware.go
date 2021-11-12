package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

var locationChan = make(chan *LocationResponse, 100)

var locationStats = map[string]int{}
var ipDictionary = map[string]*LocationResponse{}

type LocationResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func ipLocater(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestIp := c.RealIP()

		locationChan <- resolveLocation(requestIp)

		return next(c)
	}
}

func incrementCountryNum(loc *LocationResponse) {
	count, ok := locationStats[loc.Country]
	if ok {
		locationStats[loc.Country] = count + 1
	} else {
		locationStats[loc.Country] = 1
	}
}
func incrementer() {
	for loc := range locationChan {

		incrementCountryNum(loc)
	}
}

func resolveLocation(requestIp string) *LocationResponse {
	location, ok := ipDictionary[requestIp]
	//fetch from cache if exists
	if ok {
		return location
	}

	res, err := http.Get("http://ip-api.com/json/" + requestIp)
	if err != nil {
		fmt.Println(err)
	}
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var loc LocationResponse
	json.Unmarshal(respBody, &loc)
	res.Body.Close()
	fmt.Println(loc.Country, requestIp, res.StatusCode)
	//save to cache
	ipDictionary[requestIp] = &loc
	return &loc
}
