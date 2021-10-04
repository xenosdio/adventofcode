package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	adapters := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		adpt, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		adapters = append(adapters, adpt)
	}

	sort.Slice(adapters, func(i, j int) bool {
		return adapters[i] < adapters[j]
	})

	differences := map[int]int{
		1: 0,
		2: 0,
		3: 1,
	}
	differences[adapters[0]]++

	for i := 0; i < len(adapters)-1; i++ {

		diff := adapters[i+1] - adapters[i]
		differences[diff]++
	}

	prod := differences[1] * differences[3]
	log.Print(prod)
}
