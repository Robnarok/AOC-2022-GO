package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readFile()
	runPuzzleOne(input)
}

func readFile() string {
	input, error := os.ReadFile("input")
	if error != nil {
		log.Fatalf("readFile: %v", error)
	}
	return string(input)
}

func runPuzzleOne(input string) {
	status, commandtext := splitCommandsAndStatus(input)
	ship := parseStatus(status)
	commands := strings.Split(commandtext, "\n")

	for _, command := range commands {
		times, from, to := readCommand(command)
		for i := 0; i < times; i++ {
			ship = runMove(ship, from-1, to-1)
		}
	}
	printLastCargo(ship)
}

// splitCommandsAndStatus takes the Input and splits them into the Starting Status and the Command Patterns
// untested. Lets just pretend that it works
func splitCommandsAndStatus(input string) (string, string) {
	s := strings.Split(input, "\n\n")
	return s[0], s[1]
}

// This is the worst crime i could imagine
// I hate this Puzzle!
func parseStatus(input string) []string {
	lines := strings.Split(input, "\n")
	// Reverse the Lines, just to make it easier
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	var output []string
	i := 1
	for {
		index := strings.Index(lines[0], strconv.Itoa(i))
		if index == -1 {
			break
		}
		tmp := ""
		for _, v := range lines[1:] {
			if string(v[index]) == " " {
				break
			}
			tmp += string(v[index])
		}
		output = append(output, tmp)
		i++
	}
	return output
}

func printLastCargo(ship []string) {
	for _, v := range ship {
		fmt.Printf("%v", string(v[len(v)-1]))
	}
}

func runMove(ship []string, from, to int) []string {
	cargo := string(ship[from][len(ship[from])-1])
	ship[from] = strings.TrimSuffix(ship[from], cargo)
	ship[to] += cargo
	return ship
}

// input: move 1 from 2 to 1
func readCommand(command string) (int, int, int) {
	subcommand := strings.Split(command, " ")
	times, _ := strconv.Atoi(subcommand[1])
	from, _ := strconv.Atoi(subcommand[3])
	to, _ := strconv.Atoi(subcommand[5])
	return times, from, to
}
