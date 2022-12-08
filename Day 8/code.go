package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(c rune) int {
	i, _ := strconv.Atoi(string(c))
	return i
}

func findRowValueAbove(rows [][]int, row, col int) int {
	for i := row - 1; i >= 0; i-- {
		if rows[i][col] >= rows[row][col] {
			return i
		}
	}
	return -1
}

func findRowValueDown(rows [][]int, row, col int) int {
	for i := row + 1; i < len(rows); i++ {
		if rows[i][col] >= rows[row][col] {
			return i
		}
	}
	return -1
}

func findColValueLeft(rows [][]int, row, col int) int {
	for i := col - 1; i >= 0; i-- {
		if rows[row][i] >= rows[row][col] {
			return i
		}
	}
	return -1
}

func findColValueRight(rows [][]int, row, col int) int {
	for i := col + 1; i < len(rows); i++ {
		if rows[row][i] >= rows[row][col] {
			return i
		}
	}
	return -1
}

func main() {
	fileData, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	input := string(fileData)

	strRows := strings.Split(input, "\n")
	rows := [][]int{}

	for _, row := range strRows {
		newRow := []int{}
		for _, char := range row {
			newRow = append(newRow, toInt(char))
		}
		rows = append(rows, newRow)
	}

	numVisibleTrees := 0
	maxScenicScore := 0
	for rowIndex, row := range rows {
		for colIndex := range row {
			if rowIndex == 0 || rowIndex == len(rows)-1 {
				numVisibleTrees++
				continue
			}
			if colIndex == 0 || colIndex == len(row)-1 {
				numVisibleTrees++
				continue
			}

			upRow := findRowValueAbove(rows, rowIndex, colIndex)
			downRow := findRowValueDown(rows, rowIndex, colIndex)
			leftCol := findColValueLeft(rows, rowIndex, colIndex)
			rightCol := findColValueRight(rows, rowIndex, colIndex)

			if upRow == -1 || downRow == -1 || leftCol == -1 || rightCol == -1 {
				numVisibleTrees++
			}

			currentScenicScore := 1
			if upRow == -1 {
				currentScenicScore *= rowIndex
			} else {
				currentScenicScore *= rowIndex - upRow
			}

			if downRow == -1 {
				currentScenicScore *= len(rows) - 1 - rowIndex
			} else {
				currentScenicScore *= downRow - rowIndex
			}

			if leftCol == -1 {
				currentScenicScore *= colIndex
			} else {
				currentScenicScore *= colIndex - leftCol
			}

			if rightCol == -1 {
				currentScenicScore *= len(row) - 1 - colIndex
			} else {
				currentScenicScore *= rightCol - colIndex
			}

			if currentScenicScore > maxScenicScore {
				maxScenicScore = currentScenicScore
			}
		}
	}

	fmt.Printf("Answer to part 1 is %d\n", numVisibleTrees)
	// Answer to part 1 is 1647
	fmt.Printf("Answer to part 2 is %d\n", maxScenicScore)
	// Answer to part 2 is 392080
}
