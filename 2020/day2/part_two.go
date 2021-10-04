package main

import (
	"bufio"
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

	cnt := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 3 {
			log.Fatalf("error on line with content %s", line)
		}

		positions := strings.Split(parts[0], "-")
		positionA, err := strconv.Atoi(string(positions[0]))
		if err != nil {
			log.Fatal(err)
		}
		positionB, err := strconv.Atoi(string(positions[1]))
		if err != nil {
			log.Fatal(err)
		}

		policyChar := rune(parts[1][0])

		charCount := 0
		if positionA >= 1 && positionA <= len(parts[2]) {
			if rune(parts[2][positionA-1]) == policyChar {
				charCount++
			}
		}

		if positionB >= 1 && positionB <= len(parts[2]) {
			if rune(parts[2][positionB-1]) == policyChar {
				charCount++
			}
		}

		if charCount == 1 {
			cnt++
		}
	}

	log.Print(cnt)
}
