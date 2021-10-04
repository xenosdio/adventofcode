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

	myTicket := []int{}

	possibleRules := make(map[string][]bool, 0)

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

			possibleRules[ruleParts[0]] = []bool{}
		} else if inputSection == 1 {
			tickets := strings.Split(input, ",")
			for _, t := range tickets {
				ticketNum, err := strconv.Atoi(t)
				if err != nil {
					panic(err)
				}
				myTicket = append(myTicket, ticketNum)
			}

			for name := range possibleRules {
				possibleRules[name] = make([]bool, len(myTicket))
				for i := range possibleRules[name] {
					possibleRules[name][i] = true
				}
			}
		} else if inputSection == 2 {
			var notPossibleRulesLine [][]string
			var valid bool

			tickets := strings.Split(input, ",")
			for _, t := range tickets {

				valid = false
				var notPossibleRules []string

				ticketNum, err := strconv.Atoi(t)
				if err != nil {
					panic(err)
				}

				for ruleName, rule := range rules {
					if rule(ticketNum) {
						valid = true
					} else {
						notPossibleRules = append(notPossibleRules, ruleName)
					}
				}

				notPossibleRulesLine = append(notPossibleRulesLine, notPossibleRules)
				if !valid {
					break
				}
			}

			if valid {
				for i, notPossibleRules := range notPossibleRulesLine {
					for _, name := range notPossibleRules {
						possibleRules[name][i] = false
					}
				}
			}
		}
	}

	positionRules := make([]string, len(rules))

	for i := 0; i < len(positionRules); i++ {
		for name, possibilities := range possibleRules {
			var position []int
			for i, possible := range possibilities {
				if possible && positionRules[i] == "" {
					position = append(position, i)
				}
			}
			if len(position) == 1 {
				positionRules[position[0]] = name
			}
		}
	}

	product := 1

	for i, name := range positionRules {
		if strings.Fields(name)[0] == "departure" {
			product *= myTicket[i]
		}
	}

	fmt.Println(product)
}
