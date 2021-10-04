package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cnt := 0

	rules := make(map[string][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		ruleParts := strings.Split(input, " bags contain ")

		bagKey := ruleParts[0]
		rules[bagKey] = []string{}

		bagRules := strings.Split(ruleParts[1], ",")
		for _, br := range bagRules {
			words := strings.Split(strings.Trim(br, " "), " ")
			if words[0] == "no" {
				continue
			}

			bagValue := words[1] + " " + words[2]
			rules[bagKey] = append(rules[bagKey], bagValue)
		}
	}

	for ruleKey, ruleValues := range rules {
		if ruleKey == "shiny gold" {
			continue
		}

		for _, v := range ruleValues {
			if v == "shiny gold" {
				cnt++
				break
			} else if shinyGoldFound(v, rules) {
				cnt++
				break
			}
		}
	}

	log.Print(cnt)
}

func shinyGoldFound(bagKey string, rules map[string][]string) bool {

	bagRules := rules[bagKey]

	if len(bagRules) == 0 {
		return false
	}

	for _, br := range bagRules {
		if br == "shiny gold" {
			return true
		} else if shinyGoldFound(br, rules) {
			return true
		}
	}

	return false
}
