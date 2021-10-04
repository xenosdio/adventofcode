package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text()
	}

	numbers := make(map[int]int, 0)
	for i, num := range strings.Split(input, ",") {
		value, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		numbers[value] = i
	}

	var spoken int

	for i := len(numbers) + 1; i < 30000000; i++ {

		lastSpoken, seen := numbers[spoken]
		if seen {
			numbers[spoken] = i - 1
			spoken = i - 1 - lastSpoken
		} else {
			numbers[spoken] = i - 1
			spoken = 0
		}
	}

	log.Println(spoken)
}
