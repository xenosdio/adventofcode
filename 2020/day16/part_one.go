package main

import (
	"bufio"
	"fmt"
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

	rules := make(map[string]func(int) bool, 0)
	inputSection := 0

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		if input == "your ticket:" || input == "nearby tickets:" {
			inputSection++
		}
		if len(input) == 0 || input == "your ticket:" || input == "nearby tickets:" {
			continue
		}

		if inputSection == 0 {
			ruleParts := strings.Split(input, ": ")

			var min1, max1, min2, max2 int
			fmt.Sscanf(ruleParts[1], "%d-%d or %d-%d", &min1, &max1, &min2, &max2)

			rules[ruleParts[0]] = func(target int) bool {
				return (target >= min1 && target <= max1) || (target >= min2 && target <= max2)
			}
		} else if inputSection == 2 {
			tickets := strings.Split(input, ",")

			for _, t := range tickets {

				ticketNum, err := strconv.Atoi(t)
				if err != nil {
					panic(err)
				}

				valid := false

				for _, rule := range rules {
					if rule(ticketNum) {
						valid = true
						break
					}
				}

				if !valid {
					sum += ticketNum
				}
			}
		}
	}

	log.Println(sum)
}
