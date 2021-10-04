package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Operation string
	Argument  int
	Processed bool
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cnt := 0

	commands := []*Command{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			panic("wrong number of parts")
		}

		operation := parts[0]
		if operation != "jmp" && operation != "acc" && operation != "nop" {
			panic("invalid operation")
		}

		argument, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		commands = append(commands, &Command{Operation: operation, Argument: argument, Processed: false})
	}

	index := 0
	for {
		cmd := commands[index]
		if cmd.Processed {
			break
		}

		if cmd.Operation == "acc" {
			cnt += cmd.Argument
			index++
		} else if cmd.Operation == "nop" {
			index++
		} else if cmd.Operation == "jmp" {
			index += cmd.Argument
		}
		cmd.Processed = true
	}

	log.Print(cnt)
}
