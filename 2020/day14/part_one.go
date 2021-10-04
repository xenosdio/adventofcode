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

	mask := ""
	mem := make(map[string]int64)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		parts := strings.Split(scanner.Text(), " = ")

		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			number, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			value := fmt.Sprintf("%b", number)
			for len(value) != 36 {
				value = "0" + value
			}

			maskedValue := ""
			for i, bit := range mask {
				if bit != 'X' {
					maskedValue += string(bit)
				} else {
					maskedValue += string(value[i])
				}
			}

			mem[parts[0][4:len(parts[0])-1]], err = strconv.ParseInt(maskedValue, 2, 64)
			if err != nil {
				panic(err)
			}
		}
	}

	sum := 0
	for _, value := range mem {
		sum += int(value)
	}

	log.Println(sum)
}
