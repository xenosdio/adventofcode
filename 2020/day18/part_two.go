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

func solve(expression []rune, i int) (int, int) {

	var result int
	var operation rune

	var sum int
	var mul bool

	executePreviousMultiplication := func() {
		if sum > 0 {
			if mul {
				result *= sum
			} else {
				result = sum
			}
			sum = 0
		}
	}

	for ; i < len(expression); i++ {

		c := expression[i]

		var value int

		if unicode.IsDigit(c) {
			value = int(c - '0')
		} else if c == ')' {
			break
		} else if c == '(' {
			value, i = solve(expression, i+1)
		} else {
			operation = c
			continue
		}

		if operation == '*' {
			executePreviousMultiplication()
			mul = true
			sum = value
		} else {
			sum += value
		}
	}

	executePreviousMultiplication()
	return result, i
}
