package part1

import (
	"fmt"
	"math"

	"os"
	"strconv"
	"strings"
)

func Main() {
	leftSide, rightSide := formatInput(getInput())
	diffs := make([]int, len(leftSide))

	for i := range len(diffs) {
		//fmt.Println(leftSide[i], rightSide[i])
		diffs[i] = int(math.Abs(float64(leftSide[i]) - float64(rightSide[i])))

	}

	fmt.Println(add(diffs))
}

// 99974 99864
func add(input []int) int {
	sum := 0
	for i := range len(input) {
		sum += input[i]
	}
	return sum
}

func getInput() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	return string(file)
}

func formatInput(input string) ([]int, []int) {

	splitInput := strings.Split(input, "\n")

	leftSide := make([]int, len(splitInput))
	rightSide := make([]int, len(splitInput))
	for i := range len(splitInput) {
		newStringSlice := strings.Split(strings.TrimSpace(splitInput[i]), " ")
		leftSide[i], _ = strconv.Atoi(newStringSlice[0])
		rightSide[i], _ = strconv.Atoi(newStringSlice[3])
	}

	return sortInt(leftSide), sortInt(rightSide)
}

func sortInt(slice []int) []int {

	for i := 0; i < len(slice)-1; i++ {
		min := i

		for x := i + 1; x < len(slice); x++ {
			if slice[x] < slice[min] {
				min = x
			}
		}
		slice[i], slice[min] = slice[min], slice[i]

	}
	return slice
}
