package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	shipEastUnits := 0
	shipNorthUnits := 0

	waypointEastUnits := 10
	waypointNorthUnits := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		action := string(input[0])
		value, err := strconv.Atoi(input[1:])
		if err != nil {
			panic(err)
		}

		if action == "N" {
			waypointNorthUnits += value
		} else if action == "S" {
			waypointNorthUnits -= value
		} else if action == "E" {
			waypointEastUnits += value
		} else if action == "W" {
			waypointEastUnits -= value
		} else if action == "L" {
			quadrants := value / 90
			for quadrants > 0 {
				quadrants--

				waypointEastUnits, waypointNorthUnits = -1*waypointNorthUnits, waypointEastUnits
			}
		} else if action == "R" {
			quadrants := value / 90
			for quadrants > 0 {
				quadrants--

				waypointEastUnits, waypointNorthUnits = waypointNorthUnits, -1*waypointEastUnits
			}
		} else if action == "F" {
			shipEastUnits += value * waypointEastUnits
			shipNorthUnits += value * waypointNorthUnits
		} else {
			panic("invalid action")
		}
	}

	manhattan := math.Abs(float64(shipEastUnits)) + math.Abs(float64(shipNorthUnits))
	log.Println(manhattan)
}
