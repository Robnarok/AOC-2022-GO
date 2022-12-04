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
	lines := strings.Split(input, "\n")
	sum := 0
	for _, v := range lines {
		// Fix for the Newline at the End..
		// But i guess i could just remove it from the Input-File :shrug:
		if len(v) == 0 {
			continue
		}
		if challangeOne(v) {
			sum++
		}
	}
	fmt.Printf("Challange One: %v\n", sum)
	sum2 := 0
	for _, v := range lines {
		if len(v) == 0 {
			continue
		}
		if challangeTwo(v) {
			sum2++
		}
	}
	fmt.Printf("Challange two: %v\n", sum2)
}

func readFile() string {
	input, error := os.ReadFile("input")
	if error != nil {
		log.Fatalf("readFile: %v", error)
	}
	return string(input)
}

func challangeOne(input string) bool {
	ranges := strings.Split(input, ",")
	elve1, err := parseRange(ranges[0])
	if err != nil {
		log.Fatalf("challangeOne: %v", err)
	}
	elve2, err := parseRange(ranges[1])
	if err != nil {
		log.Fatalf("challangeOne: %v", err)
	}
	return checkSubspace(elve1, elve2)
}

func challangeTwo(input string) bool {
	ranges := strings.Split(input, ",")
	elve1, err := parseRange(ranges[0])
	if err != nil {
		log.Fatalf("challangeTwo: %v", err)
	}
	elve2, err := parseRange(ranges[1])
	if err != nil {
		log.Fatalf("challangeTwo: %v", err)
	}
	return findOverlaps(elve1, elve2)
}

// Returns True, if one is a subset of the other.
// i am going to hell for this implementation
func checkSubspace(elve1, elve2 []int) bool {
	// Elve 1 is a Substring of Elve2
	if elve1[0] >= elve2[0] && elve1[len(elve1)-1] <= elve2[len(elve2)-1] {
		return true
	}
	// Elve2 is a subset of Elve 1
	if elve2[0] >= elve1[0] && elve2[len(elve2)-1] <= elve1[len(elve1)-1] {
		return true
	}
	return false
}

func findOverlaps(elve1, elve2 []int) bool {
	for _, v := range elve1 {
		for _, v2 := range elve2 {
			if v == v2 {
				return true
			}
		}
	}
	return false
}

// parseRange Takes a String of a Range ("1-5") and turns it into an Array with 1,2,...5
// For Part 1 ... this is practicly useless.
func parseRange(input string) ([]int, error) {
	sectionRange := strings.Split(input, "-")

	start, err := strconv.Atoi(sectionRange[0])
	if err != nil {
		log.Fatalf("parseRange: Error Converting Start: %v", err)
		return nil, err
	}

	end, err := strconv.Atoi(sectionRange[1])
	if err != nil {
		log.Fatalf("parseRange: Error Converting End: %v", err)
		return nil, err
	}

	var section []int
	for i := start; i <= end; i++ {
		section = append(section, i)
	}
	return section, nil
}
