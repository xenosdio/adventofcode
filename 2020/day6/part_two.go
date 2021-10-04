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

	grpAnswers := make(map[rune]int, 0)
	var grpLen rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		if input == "" {
			for question, numberOfYes := range grpAnswers {
				if numberOfYes == grpAnswers[grpLen] && question != grpLen {
					cnt++
				}
			}
			grpAnswers = make(map[rune]int, 0)
		} else {
			grpAnswers[grpLen]++
			for _, r := range input {
				grpAnswers[r]++
			}
		}
	}

	for question, numberOfYes := range grpAnswers {
		if numberOfYes == grpAnswers[grpLen] && question != grpLen {
			cnt++
		}
	}

	log.Print(cnt)
}
