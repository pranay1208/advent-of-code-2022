package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleCommand(command, currentDirectory string) string {
	if strings.HasPrefix(command, "ls") {
		return currentDirectory
	}

	if command == "cd /" {
		return ""
	}

	if command == "cd .." {
		index := strings.LastIndex(currentDirectory, "/")
		return currentDirectory[:index]
	}

	return currentDirectory + "/" + command[3:]
}

func getAllDirectories(currentDirectory string) []string {
	dirs := []string{}
	dirs = append(dirs, currentDirectory)

	for strings.LastIndex(currentDirectory, "/") != -1 {
		index := strings.LastIndex(currentDirectory, "/")
		currentDirectory = currentDirectory[:index]
		dirs = append(dirs, currentDirectory)
	}

	dirs = append(dirs, "**HOME**")

	return dirs
}

func main() {
	fileData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(fileData)

	directoryToSize := make(map[string]int)

	currentDirectory := ""

	for _, line := range strings.Split(input, "\n") {
		if string(line[0]) == "$" {
			// handle as command
			_, command, _ := strings.Cut(line, " ")
			currentDirectory = handleCommand(command, currentDirectory)
			continue
		}

		// dir or number
		if strings.HasPrefix(line, "dir") {
			continue
		}

		// collect number
		numStr, _, _ := strings.Cut(line, " ")
		size, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		dirs := getAllDirectories(currentDirectory)
		for _, dir := range dirs {
			if dir == "" {
				continue
			}
			if _, found := directoryToSize[dir]; !found {
				directoryToSize[dir] = 0
			}
			directoryToSize[dir] += size
		}
	}

	answer1 := 0
	for _, dirSize := range directoryToSize {
		if dirSize < 100000 {
			answer1 += dirSize
		}
	}

	fmt.Printf("The answer to part 1 is %d\n", answer1)
	// The answer to part 1 is 1427048

	totalSize := 70000000
	usedSize, _ := directoryToSize["**HOME**"]
	unusedSpace := totalSize - usedSize
	spaceToFreeUp := 30000000 - unusedSpace

	spaceOfDirToDelete := usedSize
	dirToDelete := "/"
	for dirName, dirSize := range directoryToSize {
		if dirSize > spaceToFreeUp && dirSize < spaceOfDirToDelete {
			spaceOfDirToDelete = dirSize
			dirToDelete = dirName
		}
	}

	fmt.Printf("The answer to part 2 is %d i.e. delete %s\n", spaceOfDirToDelete, dirToDelete)
	// The answer to part 2 is 2940614 i.e. delete /tmw/gdt/snqcgbs
}
