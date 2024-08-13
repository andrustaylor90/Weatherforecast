package schema

type WeatherForecastResponse struct {
	Forecast string `json:"forecast"`
	Feeling  string `json:"feeling"`
}

type GetGridForecastURLResponse struct {
	Properties struct {
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

type GetForecastResponse struct {
	Properties struct {
		Periods []struct {
			Temperature   int32  `json:"temperature"`
			ShortForecast string `json:"shortForecast"`
		} `json:"periods"`
	} `json:"properties"`
}
