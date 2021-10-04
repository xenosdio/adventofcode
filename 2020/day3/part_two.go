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

	product := 1

	encounters := []struct {
		right int
		down  int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, e := range encounters {
		product = product * numOfTreesEncountered(input, e.right, e.down)
	}

	log.Print(product)
}

func numOfTreesEncountered(input []string, right, down int) int {

	cnt := 0

	i := 0
	rowLength := len(input[0])
	for row := 0; row < len(input); row = row + down {
		if row == 0 {
			continue
		}

		i += right
		col := i % rowLength
		if string(input[row][col]) == "#" {
			cnt++
		}
	}

	return cnt
}
