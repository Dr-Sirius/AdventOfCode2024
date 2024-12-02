package part2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Main() {

	leftSide, rightSide := formatInput(getInput())
	occurences := countOccurences(leftSide, rightSide)
	sum := 0

	for i := range len(occurences) {
		sum += (leftSide[i] * occurences[i])
	}
	fmt.Println(sum)

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

func countOccurences(left []int, right []int) []int {
	occ := make([]int, len(left))

	for i := 0; i < len(left); i++ {
		index := i
		count := 0
		for x := 0; x < len(left); x++ {
			if left[index] == right[x] {
				count++
			}
		}
		occ[i] = count

	}
	return occ
}
