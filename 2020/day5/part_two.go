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

	max := -1
	min := 2147483647

	seats := make(map[int]bool, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		row := binarySearch(input[:7], 'F', 'B', 0, 127)
		col := binarySearch(input[7:], 'L', 'R', 0, 7)

		product := (row * 8) + col
		if product > max {
			max = product
		}
		if product < min {
			min = product
		}
		seats[product] = true
	}

	for i := min; i < max; i++ {
		if !seats[i] {
			log.Print(i)
		}
	}
}

func binarySearch(input string, lowHint, highHint rune, low, high int) int {

	mid := (low + high) / 2

	for _, r := range input {

		if r == lowHint {
			high = mid
		} else if r == highHint {
			low = mid + 1
		} else {
			log.Fatalf("wrong hint: %c", r)
		}

		mid = (low + high) / 2
	}

	return mid
}
