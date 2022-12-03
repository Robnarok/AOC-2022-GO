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
	value := 0
	for _, line := range lines {
		value += paresLine(line)
	}
	fmt.Printf("value: %v\n", value)
}

func readFile() string {
	input, error := os.ReadFile("input")
	if error != nil {
		log.Fatalf("readFile: %v", error)
	}
	return string(input)
}

func paresLine(input string) int {
	s, s2 := parseItems(input)
	s3 := getCommonItem(s, s2)
	return getValue(s3)
}

func parseItems(input string) (string, string) {
	length := len(input) / 2
	return input[0:length], input[length:]
}

func getCommonItem(input1, input2 string) string {
	for _, i := range input1 {
		for _, j := range input2 {
			if i == j {
				return string(i)
			}
		}
	}
	return ""
}
func getValue(input string) int {
	value := int(input[0])
	if value > 96 {
		return value - 96
	}
	return value - 38
}
