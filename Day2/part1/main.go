package part1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Main() {

	fmt.Println(checkSafety(getDiffs(formatInput(getInput()))))
}

func getInput() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	return string(file)
}

func formatInput(input string) [][]int {
	sliceString := strings.Split(input, "\n")
	formatedInput := make([][]int, len(sliceString))

	for i := range len(sliceString) {

		splitString := strings.Split(strings.TrimSpace(sliceString[i]), " ")
		formatedInt := make([]int, len(splitString))
		for x := range len(splitString) {
			formatedInt[x], _ = strconv.Atoi(splitString[x])
		}
		formatedInput[i] = formatedInt
	}
	return formatedInput
}

func getDiffs(input [][]int) [][]int {
	diffs := make([][]int, len(input))

	for i := range len(diffs) {
		diff := make([]int, len(input[i])-1)

		for x := range len(input[i]) - 1 {
			diff[x] = input[i][x] - input[i][x+1]
		}
		diffs[i] = diff
		//fmt.Println(diff)
	}
	return diffs

}

func checkSafety(input [][]int) int {
	safety := 0

	for i := range len(input) {
		safe := true

		for x := range len(input[i]) - 1 {
			if (input[i][x] > 0 && input[i][x+1] < 0) || (input[i][x] < 0 && input[i][x+1] > 0) {
				safe = false
				break
			}
			if input[i][x] == 0 || input[i][x] > 3 || input[i][x] < -3 {
				safe = false
				break
			}
			if input[i][x+1] == 0 || input[i][x+1] > 3 || input[i][x+1] < -3 {
				safe = false
				break
			}

		}
		if safe {
			safety++
		}
		fmt.Println(input[i], safe)

	}
	return safety
}
