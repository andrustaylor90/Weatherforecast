package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Convert temperature to a subjective feeling for individuals based on temperature ranges.
Request:
  - temp: int32 temperature in Fahrenheit

Response:
  - string representing the perceived feeling associated with the given temperature.
*/
func TemperatureCharacterization(temp int32) string {
	switch {
	case temp > 86: // Hot temperature: more than 86°F (30°C)
		return "hot"
	case temp < 50: // Cold temperature: less than 50°F (10°C)
		return "cold"
	default: // Other
		return "moderate"
	}
}

/*
FetchJSON retrieves JSON data from the specified URL and parses it into the provided destination struct.
The destination can be of any type. The function returns an error if any issues occur during the process.

Request:
  - url: string representing the URL to fetch
  - dest: an interface{} to hold the parsed response

Response:
  - error: an error type indicating any issues encountered during the fetch or parsing
*/
func FetchAndDecode(url string, dest interface{}) error {
	// Send http request to get json response
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// check if response was successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error:received status code %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(dest)
}
