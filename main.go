package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// type LocationResponse struct {
// 	Status      string  `json:"status"`
// 	Country     string  `json:"country"`
// 	CountryCode string  `json:"countryCode"`
// 	Region      string  `json:"region"`
// 	RegionName  string  `json:"regionName"`
// 	City        string  `json:"city"`
// 	Zip         string  `json:"zip"`
// 	Lat         float64 `json:"lat"`
// 	Lon         float64 `json:"lon"`
// 	Timezone    string  `json:"timezone"`
// 	Isp         string  `json:"isp"`
// 	Org         string  `json:"org"`
// 	As          string  `json:"as"`
// 	Query       string  `json:"query"`
// }

//create a echo middleware
// func ipLocater(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		requestIp := c.RealIP()
// 		res, err := http.Get("http://ip-api.com/json/" + requestIp)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		respBody, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			locationStatsMutex.Lock()
// 		}
// 		var loc LocationResponse
// 		json.Unmarshal(respBody, &loc)
// 		fmt.Println(loc.Country)
// 		res.Body.Close()
// 		incrementCountryNum(&loc)
// 		return next(c)
// 	}
// }

// send a get request

func main() {
	go incrementer()
	e := echo.New()
	e.Use(ipLocater)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/stats", func(c echo.Context) error {
		return c.JSON(http.StatusOK, locationStats)
	})
	e.Logger.Fatal(e.Start("0.0.0.0:1323"))
}
