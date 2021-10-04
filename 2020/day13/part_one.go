package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) > 2 {
		panic("invalid input")
	}

	timestamp, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}

	buses := []int{}
	for _, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			continue
		}

		busID, err := strconv.Atoi(bus)
		if err != nil {
			panic(err)
		}
		buses = append(buses, busID)
	}
	sort.Ints(buses)

	minutes := -1

	i := 1
	for {
		for _, bus := range buses {
			if bus*i >= timestamp {
				minutes = (bus*i - timestamp) * bus
				break
			}
		}

		if minutes != -1 {
			break
		}

		i++
	}

	log.Println(minutes)
}
