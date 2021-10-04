package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	equations := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()
		equations = append(equations, strings.ReplaceAll(input, " ", ""))
	}

	sum := 0
	for _, equation := range equations {
		result, _ := solve([]rune(equation), 0)
		sum += result
	}

	log.Println(sum)
}

func solve(equation []rune, i int) (int, int) {

	result := 0
	var operation rune

	for ; i < len(equation); i++ {

		c := equation[i]

		var value int

		if unicode.IsDigit(c) {
			value = int(c - '0')
		} else if c == ')' {
			break
		} else if c == '(' {
			value, i = solve(equation, i+1)
		} else {
			operation = c
			continue
		}

		if operation == '*' {
			result *= value
		} else {
			result += value
		}
	}

	return result, i
}
