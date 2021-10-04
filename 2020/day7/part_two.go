package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	Name       string
	Occurences int
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cnt := 0

	rules := make(map[string][]*Bag, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		ruleParts := strings.Split(input, " bags contain ")

		bagKey := ruleParts[0]
		rules[bagKey] = []*Bag{}

		bagRules := strings.Split(ruleParts[1], ",")
		for _, br := range bagRules {
			words := strings.Split(strings.Trim(br, " "), " ")
			if words[0] == "no" {
				continue
			}

			occurences, err := strconv.Atoi(words[0])
			if err != nil {
				panic(err)
			}

			bagValue := &Bag{
				Name:       words[1] + " " + words[2],
				Occurences: occurences,
			}
			rules[bagKey] = append(rules[bagKey], bagValue)
		}
	}

	shinyGoldBags := rules["shiny gold"]
	for _, bag := range shinyGoldBags {

		cnt += bag.Occurences + (bag.Occurences * calculateInternalBags(bag.Name, rules))
	}

	log.Print(cnt)
}

func calculateInternalBags(bagKey string, rules map[string][]*Bag) int {

	bagRules := rules[bagKey]

	if len(bagRules) == 0 {
		return 0
	}

	cnt := 0

	for _, bag := range bagRules {
		cnt += bag.Occurences + (bag.Occurences * calculateInternalBags(bag.Name, rules))
	}

	return cnt
}
