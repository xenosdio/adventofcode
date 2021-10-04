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

	input := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	cnt := 0

	i := 0
	rowLength := len(input[0])
	for row := range input {
		if row == 0 {
			continue
		}

		i += 3
		col := i % rowLength
		if string(input[row][col]) == "#" {
			cnt++
		}
	}

	log.Print(cnt)
}
