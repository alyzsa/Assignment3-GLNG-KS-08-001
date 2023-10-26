package models

import (
	"database/sql"
)

type WeatherData struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

func GetLatestWeatherData(db *sql.DB) (*WeatherData, error) {
	query := "SELECT water, wind, water_status, wind_status FROM weather_data ORDER BY updated_at DESC LIMIT 1"
	row := db.QueryRow(query)

	var water, wind int
	var waterStatus, windStatus string

	err := row.Scan(&water, &wind, &waterStatus, &windStatus)
	if err != nil {
		return nil, err
	}

	latestData := &WeatherData{
		Water:       water,
		Wind:        wind,
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
	}

	return latestData, nil
}
