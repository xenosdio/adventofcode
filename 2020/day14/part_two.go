package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
			number, err := strconv.Atoi(parts[0][4 : len(parts[0])-1])
			if err != nil {
				panic(err)
			}

			address := fmt.Sprintf("%b", number)
			for len(address) != 36 {
				address = "0" + address
			}

			maskedAddress := ""
			xCnt := 0
			for i, bit := range mask {
				if bit == 'X' {
					maskedAddress += "X"
					xCnt++
				} else if bit == '1' {
					maskedAddress += "1"
				} else if bit == '0' {
					maskedAddress += string(address[i])
				} else {
					panic("invalid bit value")
				}
			}

			for i := 0; i < int(math.Pow(2, float64(xCnt))); i++ {
				combination := fmt.Sprintf("%b", i)
				for len(combination) < xCnt {
					combination = "0" + combination
				}

				tmpAddress := maskedAddress
				for _, bit := range combination {
					tmpAddress = strings.Replace(tmpAddress, "X", string(bit), 1)
				}

				value, err := strconv.Atoi(parts[1])
				if err != nil {
					panic(err)
				}
				mem[tmpAddress] = int64(value)
			}
		}
	}

	sum := 0
	for _, value := range mem {
		sum += int(value)
	}

	log.Println(sum)
}
