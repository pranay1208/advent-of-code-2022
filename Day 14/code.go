package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}

func getCoordinates(s string) (int, int) {
	parts := strings.Split(s, ",")
	xCoord, _ := strconv.Atoi(parts[0])
	yCoord, _ := strconv.Atoi(parts[1])
	return xCoord, yCoord
}

func createRockMap(input string) (map[string]bool, int) {
	rockMap := make(map[string]bool)
	maxYCoord := 0

	// part 1: Create rock map
	for _, line := range strings.Split(input, "\n") {
		// parse individual rock line
		points := strings.Split(line, " -> ")
		for index, point1 := range points {
			x1, y1 := getCoordinates(point1)
			// adjust maxYCoord
			maxYCoord = max(maxYCoord, y1)

			if index == len(points)-1 {
				continue
			}
			x2, y2 := getCoordinates(points[index+1])

			// add these points to map
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				key := fmt.Sprintf("%d|%d", x, y1)
				rockMap[key] = true
			}
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				key := fmt.Sprintf("%d|%d", x1, y)
				rockMap[key] = true
			}
		}
	}

	return rockMap, maxYCoord
}

func simulateSandFall(rockMap map[string]bool, maxY int) (bool, int, int) {
	sandX := 500
	sandY := 0

	for sandY < maxY {

		// check if floor below
		floorKey := fmt.Sprintf("FLOOR|%d", sandY+1)
		if _, found := rockMap[floorKey]; found {
			break
		}

		// try to do down
		key := fmt.Sprintf("%d|%d", sandX, sandY+1)
		if _, found := rockMap[key]; !found {
			sandY = sandY + 1
			continue
		}

		// try to go down and left
		key = fmt.Sprintf("%d|%d", sandX-1, sandY+1)
		if _, found := rockMap[key]; !found {
			sandY = sandY + 1
			sandX = sandX - 1
			continue
		}

		// try to go down and right
		key = fmt.Sprintf("%d|%d", sandX+1, sandY+1)
		if _, found := rockMap[key]; !found {
			sandY = sandY + 1
			sandX = sandX + 1
			continue
		}

		break
	}

	if sandY >= maxY {
		return true, sandX, sandY
	}

	key := fmt.Sprintf("%d|%d", sandX, sandY)
	rockMap[key] = true
	return false, sandX, sandY
}

func main() {
	fileData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(fileData)

	// Part 1
	rockMap, maxYCoord := createRockMap(input)

	count := 0

	for true {
		didInfiniteFall, _, _ := simulateSandFall(rockMap, maxYCoord)
		if didInfiniteFall {
			break
		}
		count++
	}

	fmt.Printf("Answer to part 1 is %d\n", count)
	// Answer to part 1 is 793

	// Part 2
	rockMap, maxYCoord = createRockMap(input)

	floorKey := fmt.Sprintf("FLOOR|%d", maxYCoord+2)
	rockMap[floorKey] = true

	count = 0

	for true {
		_, sandX, sandY := simulateSandFall(rockMap, maxYCoord+2)
		count++
		if sandX == 500 && sandY == 0 {
			break
		}
	}

	fmt.Printf("Answer to part 2 is %d\n", count)
	// Answer to part 2 is 24166
}
