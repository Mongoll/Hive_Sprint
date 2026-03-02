package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- Weather Station ---")

	// Map of valid sensor IDs to their keys (ordered separately)
	idToKey := map[int]string{
		1:  "airTemp",
		2:  "airPressure",
		7:  "precipitation",
		11: "windSpeed",
		12: "windDirection",
		13: "humidity",
		14: "dewPoint",
		15: "soilMoisture",
		22: "cloudCover",
	}

	// Ordered list of IDs for output
	orderedIDs := []int{1, 2, 7, 11, 12, 13, 14, 15, 22}

	// Internal state (nil means NULL)
	state := make(map[int]*float64)
	for _, id := range orderedIDs {
		state[id] = nil
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		switch {
		case line == "exit":
			fmt.Println("Exiting...")
			return

		case line == "get":
			for _, id := range orderedIDs {
				key := idToKey[id]
				if state[id] == nil {
					fmt.Printf("%s:NULL\n", key)
				} else {
					fmt.Printf("%s:%v\n", key, *state[id])
				}
			}

		case line == "clear":
			for _, id := range orderedIDs {
				state[id] = nil
			}

		default:
			// Attempt to parse id,value
			parts := strings.Split(line, ",")
			if len(parts) != 2 {
				continue // ignore invalid lines
			}

			id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				continue
			}

			// Ignore unknown IDs
			if _, exists := idToKey[id]; !exists {
				continue
			}

			valueStr := strings.TrimSpace(parts[1])

			if valueStr == "NULL" {
				state[id] = nil
				continue
			}

			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				continue
			}

			state[id] = &value
		}
	}
}