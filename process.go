package brc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Process(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	// Create a map to store temperatures for each city
	cityTemperatures := make(map[string][]float64)

	// Read lines from the input
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) != 2 {
			continue
		}

		city := parts[0]
		tempsStr := strings.Split(parts[1], "/")

		// Convert temperature strings to float and store them
		for _, tempStr := range tempsStr {
			temp, err := strconv.ParseFloat(tempStr, 64)
			if err != nil {
				log.Println("Error parsing temperature:", err)
				panic(err)
			}
			cityTemperatures[city] = append(cityTemperatures[city], temp)
		}
	}

	// Get sorted list of cities
	var cities []string
	for city := range cityTemperatures {
		cities = append(cities, city)
	}
	sort.Strings(cities)

	var finalResult string
	for _, city := range cities {
		temps := cityTemperatures[city]

		minTemp := temps[0]
		maxTemp := temps[0]
		sumTemp := 0.0

		for _, temp := range temps {
			if temp < minTemp {
				minTemp = temp
			}
			if temp > maxTemp {
				maxTemp = temp
			}
			sumTemp += temp
		}
		avgTemp := sumTemp / float64(len(temps))

		//log.Println("avg ==", avgTemp)

		finalResult += fmt.Sprintf("%s=%.2f/%.2f/%.2f\n", city, minTemp, maxTemp, avgTemp)
	}
	return finalResult, nil
}
