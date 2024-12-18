package part2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Main() {
	fInput := formatInput(getInput())

	fmt.Println(checkSafety(fInput))
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

// func getDiffs(input [][]int) [][]int {
// 	diffs := make([][]int, len(input))

// 	for i := range len(diffs) {
// 		diff := make([]int, len(input[i])-1)

// 		for x := range len(input[i]) - 1 {
// 			diff[x] = input[i][x] - input[i][x+1]
// 		}
// 		diffs[i] = diff
// 		//fmt.Println(diff)
// 	}
// 	return diffs

// }

func checkSafety(input [][]int) int {
	safety := 0

	for i := range len(input) {
		fmt.Println(input[i])
		s := checkSlice(input[i])

		if s {
			safety++
			//fmt.Println(i, input[i], s)
		} else {
			for x := 0; x < len(input[i]); x++ {
				newSlice := remove(input[i], x)
				fmt.Println(newSlice)
				slice := checkSlice(newSlice)
				if slice {
					safety++
					//fmt.Println(i, input[i], "fTs")
					break
				}

			}

		}
		fmt.Println(i, input[i], s)

	}

	return safety
}

func checkSlice(slice []int) bool {
	prevDiff := slice[0] - slice[1]
	safe := true

	for x := 1; x < len(slice)-1; x++ {
		newDiff := slice[x] - slice[x+1]
		if (newDiff > 0 && prevDiff < 0) || (newDiff < 0 && prevDiff > 0) {
			safe = false

			break
		}
		if prevDiff == 0 || prevDiff > 3 || prevDiff < -3 {
			safe = false

			break
		}
		if newDiff == 0 || newDiff > 3 || newDiff < -3 {
			safe = false

			break
		}

	}
	return safe
}

func remove(slice []int, s int) []int {

	new := make([]int, 0, len(slice)-1)
	new = append(new, slice[:s]...)
	new = append(new, slice[s+1:]...)

	return new
}
