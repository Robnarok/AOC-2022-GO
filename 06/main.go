package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input := readFile()
	fmt.Printf("searchSequence(input): %v\n", searchSequence(input))
}

func readFile() string {
	input, error := os.ReadFile("input")
	if error != nil {
		log.Fatalf("readFile: %v", error)
	}
	return string(input)
}

func searchSequence(input string) int {

	lenght := len(input)

	for i := 3; i < lenght; i++ {
		marktext := collectSequenceFromPos(input, i)
		if isSignal(marktext) {
			return i + 1
		}
	}

	return -1
}

func collectSequenceFromPos(input string, place int) string {
	if place < 3 {
		return ""
	}
	returnString := ""
	returnString += string(input[place-3])
	returnString += string(input[place-2])
	returnString += string(input[place-1])
	returnString += string(input[place])
	return returnString
}

// isSignal gets a String with 4 Chars and checks if there is a double Occurence
func isSignal(input string) bool {
	length := len(input)
	if length != 4 {
		log.Fatalf("isSignal: %v", "Input is to long")
	}
	// Only needs to check 3 Chars
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}
	return true
}
