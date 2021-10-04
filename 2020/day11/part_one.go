package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seats := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		rowSeats := []string{}
		for _, r := range input {
			rowSeats = append(rowSeats, string(r))
		}

		seats = append(seats, rowSeats)
	}

	var occupied int

	for {
		changes := 0

		newSeats := [][]string{}
		for i := 0; i < len(seats); i++ {
			newRowSeats := []string{}

			for j := 0; j < len(seats[i]); j++ {
				if seats[i][j] == "." {
					newRowSeats = append(newRowSeats, ".")
				} else if seats[i][j] == "L" {
					foundOccupiedSeat := false

					for x := -1; x <= 1; x++ {
						for y := -1; y <= 1; y++ {
							if i+y >= len(seats) || j+x >= len(seats[i]) || i+y < 0 || j+x < 0 {
								continue
							}

							if (x != 0 || y != 0) && seats[i+y][j+x] == "#" {
								foundOccupiedSeat = true
							}
						}
					}

					if !foundOccupiedSeat {
						newRowSeats = append(newRowSeats, "#")
						changes++
					} else {
						newRowSeats = append(newRowSeats, "L")
					}
				} else { // seats[i][j] == "#"
					occupiedSeats := 0

					for x := -1; x <= 1; x++ {
						for y := -1; y <= 1; y++ {
							if i+y >= len(seats) || j+x >= len(seats[i]) || i+y < 0 || j+x < 0 {
								continue
							}

							if (x != 0 || y != 0) && seats[i+y][j+x] == "#" {
								occupiedSeats++
							}
						}
					}

					if occupiedSeats >= 4 {
						newRowSeats = append(newRowSeats, "L")
						changes++
					} else {
						newRowSeats = append(newRowSeats, "#")
					}
				}
			}

			newSeats = append(newSeats, newRowSeats)
		}

		seats = newSeats

		if changes == 0 {
			for i := 0; i < len(seats); i++ {
				for j := 0; j < len(seats[i]); j++ {
					if seats[i][j] == "#" {
						occupied++
					}
				}
			}
			break
		}
	}

	log.Print(occupied)
}
