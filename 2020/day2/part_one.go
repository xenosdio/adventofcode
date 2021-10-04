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

		timeParts := strings.Split(parts[0], "-")
		minTimes, err := strconv.Atoi(string(timeParts[0]))
		if err != nil {
			log.Fatal(err)
		}
		maxTimes, err := strconv.Atoi(string(timeParts[1]))
		if err != nil {
			log.Fatal(err)
		}

		policyChar := rune(parts[1][0])

		charCount := 0
		for _, letter := range parts[2] {
			if letter == policyChar {
				charCount++
			}
		}

		if charCount >= minTimes && charCount <= maxTimes {
			cnt++
		}
	}

	log.Print(cnt)
}
