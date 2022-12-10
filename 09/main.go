package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello THere")
	lines := strings.Split(readFile(), "\n")
	head := position{0, 0}
	tail := position{0, 0}
	field := map[position]bool{
		tail: true,
	}
	for _, v := range lines {
		head, tail, field = runCommand(head, tail, v, field)
	}
	fmt.Printf("len(field): %v\n", len(field))
}

func readFile() string {
	input, error := os.ReadFile("input")
	if error != nil {
		log.Fatalf("readFile: %v", error)
	}
	return string(input)
}
func runCommand(head, tail position, command string, fields map[position]bool) (position, position, map[position]bool) {
	s := strings.Split(command, " ")
	count, _ := strconv.Atoi(s[1])
	for j := 0; j < count; j++ {
		head, tail = move(head, tail, s[0])
		fields[tail] = true
	}
	return head, tail, fields
}

func move(head, tail position, command string) (position, position) {
	previus := head
	switch command {
	case "R":
		head.moveRight()
	case "L":
		head.moveLeft()
	case "U":
		head.moveUp()
	case "D":
		head.moveDown()
	default:
	}
	if checkDetached(head, tail) {
		return head, previus
	}
	return head, tail
}

func checkDetached(head, tail position) bool {
	// Check if X is to far away
	if head.posx+2 == tail.posx {
		return true
	}
	if head.posx-2 == tail.posx {
		return true
	}
	// check if Y is to far away
	if head.posy+2 == tail.posy {
		return true
	}
	if head.posy-2 == tail.posy {
		return true
	}
	return false
}

type position struct {
	posx, posy int
}

func (position *position) moveRight() {
	position.posx += 1
}

func (position *position) moveLeft() {
	position.posx -= 1
}
func (position *position) moveUp() {
	position.posy += 1
}

func (position *position) moveDown() {
	position.posy -= 1
}
