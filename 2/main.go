package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input := readFile()
	lines := strings.Split(input, "\n")

	sum := 0
	for _, l := range lines {
		if len(l) == 3 {
			sum += playGame(l)
		}
	}
	fmt.Printf("sum: %v\n", sum)
}

func readFile() string {
	input, error := os.ReadFile("input")
	if error != nil {
		log.Fatalf("readFile: %v", error)
	}
	return string(input)
}

func playGame(input string) int {
	sum := 0
	game := strings.Split(input, " ")
	sum += checkWin(game[0], game[1])
	sum += scoreChoice(game[1])
	return sum
}

func scoreChoice(you string) int {
	if you == "X" {
		return 1
	}
	if you == "Y" {
		return 2
	}
	return 3
}

// Convert the Input to Int and Compare, 23 is the Offset
// I am really not that happy with the way i check if i've won
func checkWin(enemy, you string) int {
	// Draw Case
	if enemy[0]+23 == you[0] {
		return 3
	}
	// Win Case
	if enemy == "A" && you == "Y" {
		return 6
	}

	if enemy == "B" && you == "Z" {
		return 6
	}

	if enemy == "C" && you == "X" {
		return 6
	}
	// Loose Case
	return 0
}
