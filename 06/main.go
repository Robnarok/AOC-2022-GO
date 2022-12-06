package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input := readFile()
	fmt.Printf("Part One: %v\n", searchSequence(4, input))
	fmt.Printf("Part Two: %v\n", searchSequence(14, input))
}

func readFile() string {
	input, error := os.ReadFile("input")
	if error != nil {
		log.Fatalf("readFile: %v", error)
	}
	return string(input)
}

func searchSequence(sequenceLength int, input string) int {

	lenght := len(input)

	for i := sequenceLength - 1; i < lenght; i++ {
		marktext := collectSequenceFromPos(sequenceLength, input, i)
		if isSignal(sequenceLength, marktext) {
			return i + 1
		}
	}

	return -1
}

func collectSequenceFromPos(sequenceLength int, input string, place int) string {
	if place < sequenceLength-1 {
		return ""
	}
	returnString := ""
	for i := sequenceLength - 1; i >= 0; i-- {
		returnString += string(input[place-i])
	}
	return returnString
}

// isSignal gets a String with 4 Chars and checks if there is a double Occurence
func isSignal(sequencelength int, input string) bool {
	length := len(input)
	if length != sequencelength {
		log.Fatalf("isSignal: %v", "Input is not the right size")
	}
	// The last Char doesnt need to be checked
	for i := 0; i < sequencelength-1; i++ {
		for j := i + 1; j < sequencelength; j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}
	return true
}
