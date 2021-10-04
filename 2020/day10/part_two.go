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

	adapters = append([]int{0}, adapters...)
	cnt := solutionsFromIndex(adapters, 0, make(map[int]int))

	log.Print(cnt)
}

func solutionsFromIndex(adapters []int, i int, cache map[int]int) int {
	if n, ok := cache[i]; ok {
		return n
	}

	if i == len(adapters)-1 {
		return 1
	}

	var solutions int
	for j := i + 1; j < len(adapters); j++ {
		if diff := adapters[j] - adapters[i]; diff > 3 {
			break
		}

		solutions += solutionsFromIndex(adapters, j, cache)
	}

	cache[i] = solutions
	return solutions
}
