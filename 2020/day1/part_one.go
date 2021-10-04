package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, i)
	}

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				log.Print(input[i] * input[j])
				break
			}
		}
	}
}
