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
							for distance := 1; distance < len(seats); distance++ {

								if i+y*distance >= len(seats) || j+x*distance >= len(seats[i]) || i+y*distance < 0 || j+x*distance < 0 {
									continue
								}

								if (x != 0 || y != 0) && seats[i+y*distance][j+x*distance] == "#" {
									foundOccupiedSeat = true
									break
								} else if (x != 0 || y != 0) && seats[i+y*distance][j+x*distance] == "L" {
									break
								}
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
							for distance := 1; distance < len(seats); distance++ {

								if i+y*distance >= len(seats) || j+x*distance >= len(seats[i]) || i+y*distance < 0 || j+x*distance < 0 {
									continue
								}

								if (x != 0 || y != 0) && seats[i+y*distance][j+x*distance] == "#" {
									occupiedSeats++
									break
								} else if (x != 0 || y != 0) && seats[i+y*distance][j+x*distance] == "L" {
									break
								}
							}
						}
					}

					if occupiedSeats >= 5 {
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
