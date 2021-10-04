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

	eastUnits := 0
	northUnits := 0

	face := "east"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		action := string(input[0])
		value, err := strconv.Atoi(input[1:])
		if err != nil {
			panic(err)
		}

		if action == "N" {
			northUnits += value
		} else if action == "S" {
			northUnits -= value
		} else if action == "E" {
			eastUnits += value
		} else if action == "W" {
			eastUnits -= value
		} else if action == "L" {
			quadrants := value / 90
			for quadrants > 0 {
				quadrants--
				if face == "east" {
					face = "north"
				} else if face == "west" {
					face = "south"
				} else if face == "north" {
					face = "west"
				} else { // south
					face = "east"
				}
			}
		} else if action == "R" {
			quadrants := value / 90
			for quadrants > 0 {
				quadrants--
				if face == "east" {
					face = "south"
				} else if face == "west" {
					face = "north"
				} else if face == "north" {
					face = "east"
				} else { // south
					face = "west"
				}
			}
		} else if action == "F" {
			if face == "east" {
				eastUnits += value
			} else if face == "west" {
				eastUnits -= value
			} else if face == "north" {
				northUnits += value
			} else { // south
				northUnits -= value
			}
		} else {
			panic("invalid action")
		}
	}

	manhattan := math.Abs(float64(eastUnits)) + math.Abs(float64(northUnits))
	log.Println(manhattan)
}
