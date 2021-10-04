package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var invalidNum int

	index := 0
	preamble := 25

	nums := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		num, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}

		if index < preamble {
			nums = append(nums, num)
			index++
			continue
		}

		if !isNumValid(num, preamble, nums) {
			invalidNum = num
			break
		}

		nums = append(nums, num)
		index++
	}

	log.Print(invalidNum)
}

func isNumValid(candidate, preamble int, nums []int) bool {

	for i := 0; i < preamble; i++ {
		for j := i + 1; j < preamble; j++ {
			if nums[len(nums)-1-i]+nums[len(nums)-1-j] == candidate {
				return true
			}
		}
	}

	return false
}
