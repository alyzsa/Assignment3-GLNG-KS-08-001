package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/alyzsa/Assignment3-GLNG-KS-08-001/database"
	"github.com/alyzsa/Assignment3-GLNG-KS-08-001/handlers"
)

func main() {
	if err := database.InitializeDatabase(); err != nil {
		log.Fatal(err)
	}
	defer database.CloseDatabase()
	router := handlers.CreateRouter()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("The server is running on :8080")
	go updateDataEvery15Seconds()
	log.Fatal(server.ListenAndServe())
}

func updateDataEvery15Seconds() {
	for {
		err := generateAndSaveData()
		if err != nil {
			log.Println("Failed to update data", err)
		}
		time.Sleep(15 * time.Second)
	}
}

func generateAndSaveData() error {
	rand.Seed(time.Now().UnixNano())
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	var waterStatus string
	if water < 5 {
		waterStatus = "safe"
	} else if water >= 5 && water <= 8 {
		waterStatus = "alert"
	} else {
		waterStatus = "danger"
	}

	var windStatus string
	if wind < 5 {
		windStatus = "safe"
	} else if wind >= 5 && wind <= 8 {
		windStatus = "alert"
	} else {
		windStatus = "danger"
	}

	_, err := database.DB.Exec("INSERT INTO weather_data (water, wind, water_status, wind_status) VALUES (?, ?, ?, ?)", water, wind, waterStatus, windStatus)
	if err != nil {
		return err
	}

	log.Printf("water: %d, wind: %d\n", water, wind)
	log.Printf("water status: %s, wind status: %s\n", waterStatus, windStatus)

	return nil
}
