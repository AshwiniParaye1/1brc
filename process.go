package brc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func Process(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)

	cityTemperatures := make(map[string]struct {
		min, max, sum float64
		count         int
	})

	// Read lines from the input
	for scanner.Scan() {
		line := scanner.Bytes()
		parts := bytes.SplitN(line, []byte(";"), 2) // Split only once to handle semicolons in city names
		if len(parts) != 2 {
			continue
		}
		city := string(parts[0])
		tempsStr := parts[1]

		// Convert temperature strings to float and update city statistics
		for _, tempStr := range bytes.Split(tempsStr, []byte("/")) {
			temp, err := strconv.ParseFloat(string(tempStr), 64)
			if err != nil {
				return "", fmt.Errorf("error parsing temperature: %v", err)
			}
			stats, ok := cityTemperatures[city]
			if !ok {
				stats.min, stats.max = temp, temp
			} else {
				if temp < stats.min {
					stats.min = temp
				}
				if temp > stats.max {
					stats.max = temp
				}
			}
			stats.sum += temp
			stats.count++
			cityTemperatures[city] = stats
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Get sorted list of cities
	cities := make([]string, 0, len(cityTemperatures))
	for city := range cityTemperatures {
		cities = append(cities, city)
	}
	sort.Strings(cities)

	var buf bytes.Buffer

	// Process temperatures for each city and append result to buf
	for _, city := range cities {
		stats := cityTemperatures[city]
		avg := stats.sum / float64(stats.count)
		buf.WriteString(fmt.Sprintf("%s=%.2f/%.2f/%.2f\n", city, stats.min, stats.max, avg))
	}

	return buf.String(), nil
}
