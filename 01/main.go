package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := readFile()
	elves := parseElves(input)
	var sums []int
	for _, elf := range elves {
		sums = append(sums, elf.sum())
	}

	sort.Ints(sums)
	fmt.Printf("result Part 1: %v\n", sums[len(sums)-1])
	fmt.Printf("result Part 2: %v\n", sums[len(sums)-1]+sums[len(sums)-2]+sums[len(sums)-3])
}

func readFile() string {
	input, error := os.ReadFile("input/input")
	if error != nil {
		log.Fatalf("readFile: %v", error)
	}
	return string(input)
}

func parseElves(input string) []Elf {
	strArr := strings.Split(input, "\n")
	var elves []Elf
	var tmp Elf
	for _, entry := range strArr {
		if entry == "" {
			//fmt.Printf("tmp: %v\n", tmp)
			elves = append(elves, tmp)
			tmp = Elf{}
			continue
		}
		i, _ := strconv.Atoi(entry)
		tmp.calorins = append(tmp.calorins, i)
	}
	return elves
}

type Elf struct {
	calorins []int
}

func (elf Elf) sum() int {
	result := 0

	for _, num := range elf.calorins {
		result += num
	}
	return result
}
