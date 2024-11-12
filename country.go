package main

import (
	"fmt"
	"math"
)

type Country struct {
	Name      string
	Latitude  float64
	Longitude float64
}

// Haversine formula to calculate distance in kilometers
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	// Radius of Earth in kilometers
	const radius = 6371.0

	// Convert degrees to radians
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	// Differences in coordinates
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Haversine formula
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	distance := radius * c
	return distance
}

func main() {
	// List of countries and their coordinates (lat, lon)
	countries := []Country{
		{"USA", 37.7749, -122.4194},       // San Francisco, USA
		{"Germany", 51.1657, 10.4515},     // Germany
		{"Australia", -25.2744, 133.7751}, // Australia
		{"India", 20.5937, 78.9629},       // India
		{"Brazil", -14.2350, -51.9253},    // Brazil
	}

	// Loop through each pair of countries and calculate distance
	for i := 0; i < len(countries); i++ {
		for j := i + 1; j < len(countries); j++ {
			country1 := countries[i]
			country2 := countries[j]
			distance := haversine(country1.Latitude, country1.Longitude, country2.Latitude, country2.Longitude)
			fmt.Printf("Distance between %s and %s: %.2f km\n", country1.Name, country2.Name, distance)
		}
	}
}
