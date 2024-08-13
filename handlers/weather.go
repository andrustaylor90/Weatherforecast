package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PowerLightStar/WeatherForecast/schema"
	"github.com/PowerLightStar/WeatherForecast/utils"
)

/*
GET: WeatherHandler retrieves the weather forecast based on the specified latitude and longitude.

Parameters:
  - longitude: The longitude coordinate of the location.
  - latitude: The latitude coordinate of the location.
*/
func WeatherHandler(w http.ResponseWriter, r *http.Request) {

	// Parse latitude and longitude from the request parameters.
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	// make a url to access weather api server
	url := fmt.Sprintf("https://api.weather.gov/points/%s,%s", lat, lon)

	w.Header().Set("Content-Type", "application/json")

	// get grid data from latitude and longitude of point
	var jsonResp schema.GetGridForecastURLResponse
	if err := utils.FetchAndDecode(url, &jsonResp); err != nil {
		print(err.Error())
		http.Error(w, "Failed to decode location data", http.StatusInternalServerError)
		// json.NewEncoder(w).Encode({"message": "Failed to decode location data"})
		return
	}

	// get forecast response data
	var jsonData schema.GetForecastResponse
	if err := utils.FetchAndDecode(jsonResp.Properties.Forecast, &jsonData); err != nil {
		// json.NewEncoder(w).Encode({"message": "Failed to get forecast data"})
		http.Error(w, "Failed to get forecast data", http.StatusInternalServerError)
		return
	}

	// check if there is no forecast data
	if len(jsonData.Properties.Periods) == 0 {
		http.Error(w, "No forecast available", http.StatusNotFound)
		return
	}

	// Get relevant data for response
	weatherToday := jsonData.Properties.Periods[0].ShortForecast
	tempToday := jsonData.Properties.Periods[0].Temperature
	// parse the tempeature to feeling
	feelingToday := utils.TemperatureCharacterization(tempToday)

	// Make a response for this handler
	weatherResp := schema.WeatherForecastResponse{
		Forecast: weatherToday,
		Feeling:  feelingToday,
	}
	json.NewEncoder(w).Encode(weatherResp)
}
