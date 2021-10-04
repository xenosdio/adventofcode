package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/deanveloper/modmath/bigmod"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) > 2 {
		panic("invalid input")
	}

	buses := []int{}
	for _, bus := range strings.Split(lines[1], ",") {
		if bus == "x" {
			buses = append(buses, -1)
			continue
		}

		busID, err := strconv.Atoi(bus)
		if err != nil {
			panic(err)
		}
		buses = append(buses, busID)
	}

	timestamp, err := solveCrt(buses)
	if err != nil {
		panic(err)
	}
	log.Println(timestamp)
}

func solveCrt(busIDs []int) (t *big.Int, err error) {
	// The bigmod library uses panics
	defer func() {
		if r := recover(); r != nil && err == nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()

	entries := make([]bigmod.CrtEntry, 0, len(busIDs))
	for i, busID := range busIDs {
		if busID == -1 {
			continue
		}

		entries = append(entries, bigmod.CrtEntry{
			A: big.NewInt(int64(-i)),
			N: big.NewInt(int64(busID)),
		})
	}
	return bigmod.SolveCrtMany(entries), nil
}
