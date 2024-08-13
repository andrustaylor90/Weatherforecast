package routes

import (
	"net/http"

	"github.com/PowerLightStar/WeatherForecast/handlers"
)

/*
registerRoutes manages all endpoints and their corresponding handler functions.
*/
func RegisterRoutes() {
	http.HandleFunc("/weather", handlers.WeatherHandler)
}
