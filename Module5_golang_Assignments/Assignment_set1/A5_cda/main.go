package main

import (
	"errors"
	"fmt"
	"strings"
)

type City struct {
	Name           string
	Temperature float64
	Rainfall       float64
}

//fuction to display cities
func displayCities(cities []City) {
		fmt.Println("\nCity Data:")
		for _, city := range cities {
			fmt.Printf(" %s:\n Temperature : %.2f°C, Rainfall : %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
		}
	}

 // Function to find the city with the highest and lowest temperatures
func findTemperatureExtremes(cities []City) (highest, lowest City) {
	highest, lowest = cities[0], cities[0]
	for _, city := range cities {
		if city.Temperature > highest.Temperature {
			highest = city
		}
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return
}


func averageRainfall(cities []City) float64 {
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

func filterCitiesByRainfall(cities []City, threshold float64) []City {
	filteredCities := []City{}
	for _, city := range cities {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities
}

func searchCityByName(cities []City, name string) (*City, error) {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return &city, nil
		}
	}
	return nil, errors.New("city not found")
}

func main() {
	cities := []City{
		{"New York", 12.5, 1200.0},
		{"Los Angeles", 18.7, 380.0},
		{"Chicago", 10.2, 990.0},
		{"America", 20.3, 1310.0},
		{"Canada", 22.8, 230.0},
	}

	fmt.Println("Climate Data Analysis")
	fmt.Println("1. Display All Cities")
	fmt.Println("2. Display city with the Highest and lowest temperature")
	fmt.Println("3. Calculate average rainfall")
	fmt.Println("4. Filter cities by rainfall")
	fmt.Println("5. Search for a city by name")
	for {
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
				displayCities(cities)


		case 2:
			highest, lowest := findTemperatureExtremes(cities)
      		fmt.Printf("City with the Highest Temperature: %s (%.2f°C)\n", highest.Name, highest.Temperature)
			fmt.Printf("City with the Lowest Temperature: %s (%.2f°C)\n", lowest.Name, lowest.Temperature)

		case 3:
			average := averageRainfall(cities)
			fmt.Printf("Average rainfall: %.2f mm\n", average)

		case 4:
			var threshold float64
			fmt.Print("Enter rainfall threshold: ")
			fmt.Scan(&threshold)
			filteredCities := filterCitiesByRainfall(cities, threshold)
			if len(filteredCities) == 0 {
				fmt.Println("No cities with rainfall above the threshold.")
			} else {
				fmt.Println("Cities with rainfall above the threshold:")
				for _, city := range filteredCities {
					fmt.Printf("%s: %.2f mm\n", city.Name, city.Rainfall)
				}
			}

		case 5:
			var name string
			fmt.Print("Enter city name: ")
			fmt.Scan(&name)
			if city, err := searchCityByName(cities, name); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("City found: %s,  Temperature: %.2f°C, Rainfall: %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
			}

		default:
			fmt.Println("Exiting system. Goodbye!")
			return
		}
	}
}