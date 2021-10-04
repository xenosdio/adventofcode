package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cnt := 0

	grpAnswers := make(map[rune]bool, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		if len(input) == 0 {
			cnt += len(grpAnswers)
			grpAnswers = make(map[rune]bool, 0)
		}

		for _, r := range input {
			grpAnswers[r] = true
		}
	}

	cnt += len(grpAnswers)
	log.Print(cnt)
}
