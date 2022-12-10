package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type folder struct {
	dirPath string
	folders []string
	files   []int
}

func (folder folder) sum(computer map[string]folder) int {
	sum := 0
	for _, v := range folder.files {
		sum += v
	}
	for _, v := range folder.folders {
		sum += computer[v].sum(computer)
	}
	return sum
}

func main() {
	input := readFile()
	dir := "/"
	computer := map[string]folder{
		"/": {
			dirPath: "/",
			folders: []string{},
			files:   []int{},
		},
	}
	lines := strings.Split(input, "\n")
	for _, v := range lines {
		dir, computer = runLine(dir, v, computer)
	}
	sum := 0
	for _, f := range computer {
		tmp := f.sum(computer)
		if tmp <= 100000 {
			sum += tmp
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

func runLine(dir, command string, computer map[string]folder) (string, map[string]folder) {
	commandPart := strings.Split(command, " ")
	if commandPart[0] == "$" {
		if commandPart[1] == "ls" {
			return dir, computer
		}
		if commandPart[1] == "cd" {
			s := changeDir(dir, commandPart[2])
			return s, computer
		}
	}
	if commandPart[0] == "dir" {
		m := parseFolders(dir, commandPart[1], computer)
		return dir, m
	}
	m := parseFiles(dir, commandPart[0], computer)
	return dir, m
}

func parseFolders(dir, command string, computer map[string]folder) map[string]folder {
	currentFolder := computer[dir]
	folderPath := dir + "/" + command
	currentFolder.folders = append(currentFolder.folders, folderPath)
	computer[folderPath] = folder{
		dirPath: folderPath,
		folders: []string{},
		files:   []int{},
	}
	computer[dir] = currentFolder
	return computer

}

func parseFiles(dir, command string, computer map[string]folder) map[string]folder {
	currentFolder := computer[dir]
	size, _ := strconv.Atoi(command)
	currentFolder.files = append(currentFolder.files, size)
	computer[dir] = currentFolder
	return computer
}

func changeDir(current, dir string) string {
	if dir == "/" {
		return "/"
	}
	if dir == ".." {
		if current == "/" {
			return "/"
		}
		dirs := strings.Split(current, "/")
		if len(dirs) == 2 { // If the Path is only /foo/ and we cd .. we want / not empty
			return "/"
		}
		current = ""
		for i, folder := range dirs {
			if folder == "" {
				continue
			}
			if i == len(dirs)-1 {
				return current
			}
			current += "/" + folder

		}

		return current
	}
	if current == "/" {
		return "/" + dir
	}

	return current + "/" + dir
}
