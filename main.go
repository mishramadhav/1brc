package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	measurementFilePath = "/Users/mukund.madhav/open-source/1brc/measurements.txt"
	semicolon           = ";"
)

type result struct {
	min, max, total float64
	count           int
}

func main() {
	resultsByCity := make(map[string]*result)

	// Read the file
	// For each line in the file
	// Parse the line
	// If the city is not in the map, add it
	// Update the result for the city
	file, e := os.OpenFile(measurementFilePath, os.O_RDONLY, 0)
	if e != nil {
		fmt.Println(e)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, _, e := reader.ReadLine()
		if e == io.EOF {
			break
		}

		// Parse the line
		splitString := strings.Split(string(line), semicolon)
		city, temperature := splitString[0], splitString[1]

		// If the city is not in the map, add it
		if _, ok := resultsByCity[city]; !ok {
			resultsByCity[city] = &result{}
		}

		// Update the result for the city
		temperature = strings.TrimSpace(temperature)
		parsedTemperature, e := strconv.ParseFloat(temperature, 64)
		if e != nil {
			fmt.Println(e)
			return
		}

		// make temperature int by multiplying by 10
		parsedTemperature = parsedTemperature * 10
		parsedTemperature = math.Trunc(parsedTemperature)

		currentResult := resultsByCity[city]
		currentResult.count++
		currentResult.total += parsedTemperature
		if currentResult.min == 0 || currentResult.min > parsedTemperature {
			currentResult.min = parsedTemperature
		}
		if currentResult.max == 0 || currentResult.max < parsedTemperature {
			currentResult.max = parsedTemperature
		}
	}

	// Print the results
	for city, result := range resultsByCity {
		fmt.Printf("%s=%f/%f/%f\n", city, result.min/10, result.total/float64(result.count)/10, result.max/10)
	}
}
