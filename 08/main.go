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
	forest := parseInputToForest(input)
	fmt.Printf("countVisible(forest): %v\n", countVisible(forest))
}

func readFile() string {
	input, error := os.ReadFile("input")
	if error != nil {
		log.Fatalf("readFile: %v", error)
	}
	return string(input)
}

func parseInputToForest(input string) [][]int {
	lines := strings.Split(input, "\n")
	var forest [][]int
	for _, v := range lines {
		row := []int{}
		s := strings.Split(v, "")
		for _, v2 := range s {
			tmp, _ := strconv.Atoi(v2)
			row = append(row, tmp)
		}
		forest = append(forest, row)
	}
	return forest
}

func countVisible(forest [][]int) int {
	counter := 0
	for y, v := range forest {
		for x, _ := range v {
			if checkVisibility(forest, x, y) {
				counter++
			}
		}
	}
	return counter
}

func checkVisibility(forest [][]int, posx, posy int) bool {
	if checkVisibilityUp(forest, posx, posy) {
		return true
	}
	if checkVisibilityDown(forest, posx, posy) {
		return true
	}
	if checkVisibilityRight(forest, posx, posy) {
		return true
	}
	if checkVisibilityLeft(forest, posx, posy) {
		return true
	}
	return false
}

// Returns True, if the Tree is Visible from the Right (Higher then all Trees)
func checkVisibilityRight(forest [][]int, posx, posy int) bool {
	if posx+1 == len(forest[posy]) {
		return true
	}

	// XI
	for i := posx + 1; i < len(forest[posy]); i++ {
		if forest[posy][posx] <= forest[posy][i] {
			return false
		}
	}
	return true
}

// Returns True, if the Tree is Visible from the left (Higher then all Trees)
func checkVisibilityLeft(forest [][]int, posx, posy int) bool {
	if posx == 0 {
		return true
	}

	for i := 0; i < posx; i++ {
		if forest[posy][posx] <= forest[posy][i] {
			return false
		}
	}
	return true
}

// Returns True, if the Tree is Visible from upwards (Higher then all Trees)
func checkVisibilityUp(forest [][]int, posx, posy int) bool {
	if posy == 0 {
		return true
	}

	for i := 0; i < posy; i++ {
		if forest[posy][posx] <= forest[i][posx] {
			return false
		}
	}
	return true
}

// Returns True, if the Tree is Visible from downwards(Higher then all Trees)
func checkVisibilityDown(forest [][]int, posx, posy int) bool {
	if posy+1 == len(forest) {
		return true
	}

	for i := posy + 1; i < len(forest); i++ {
		if forest[posy][posx] <= forest[i][posx] {
			return false
		}
	}
	return true
}
