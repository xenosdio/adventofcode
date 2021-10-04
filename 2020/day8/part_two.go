package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	minInt = -2147483648
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

	for i := 0; i < len(commands); i++ {
		if commands[i].Operation == "acc" {
			continue
		} else if commands[i].Operation == "jmp" {
			newCommands := []*Command{}
			for _, cmd := range commands {
				newCommands = append(newCommands, &Command{Operation: cmd.Operation, Argument: cmd.Argument, Processed: cmd.Processed})
			}

			newCommands[i].Operation = "nop"
			cnt = runProgram(newCommands)
			if cnt != minInt {
				break
			}
		} else if commands[i].Operation == "nop" {
			newCommands := []*Command{}
			for _, cmd := range commands {
				newCommands = append(newCommands, &Command{Operation: cmd.Operation, Argument: cmd.Argument, Processed: cmd.Processed})
			}

			newCommands[i].Operation = "jmp"
			cnt = runProgram(newCommands)
			if cnt != minInt {
				break
			}
		}
	}

	log.Print(cnt)
}

func runProgram(commands []*Command) int {
	cnt := 0

	index := 0
	for {
		if index > len(commands)-1 {
			break
		}

		cmd := commands[index]
		if cmd.Processed {
			cnt = minInt
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

	return cnt
}
