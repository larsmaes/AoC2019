package main

import (
	"fmt"
	"strconv"
)

func generateInputs(from int, to int) []int {

	var inputs []int
	for i := from; i <= to; i++ {
		inputs = append(inputs, i)
	}

	return inputs
}

func getPasswordsWithAdjacentDigits(inputs []int) (output []int) {

	for _, input := range inputs {
		if hasAdjacentDigits(input) {
			output = append(output, input)
		}
	}

	return
}

func hasAdjacentDigits(check int) bool {
	hasAdjacent := false

	digitString := strconv.Itoa(check)

	for k, v := range digitString {
		if k == 0 {
			continue
		}
		if string(v) == digitString[k-1:k] {
			hasAdjacent = true
		}
	}

	return hasAdjacent
}

func hasNoDecreasingDigits(check int) bool {
	hasNoDecreasing := true

	digitString := strconv.Itoa(check)

	for k, v := range digitString {
		if k == 0 {
			continue
		}

		a, _ := strconv.Atoi(string(v))
		b, _ := strconv.Atoi(digitString[k-1 : k])

		if a < b {
			hasNoDecreasing = false
		}
	}

	return hasNoDecreasing
}

func getPasswordWithoutDecreasingDigits(inputs []int) (output []int) {
	for _, input := range inputs {
		if hasNoDecreasingDigits(input) {
			output = append(output, input)
		}
	}

	return
}

func getPasswordsNoPartOfLargerGroups(inputs []int) (output []int) {
	for _, input := range inputs {
		if hasNoLargerPartDigits(input) {
			output = append(output, input)
		}
	}

	return
}

func hasNoLargerPartDigits(check int) bool {

	digitString := strconv.Itoa(check)

	seen := make(map[string]int)

	for _, v := range digitString {
		seen[string(v)]++
	}

	for _, v := range seen {
		if v == 2 {
			return true
		}
	}

	return false
}

func getAmountOfMatchingPasswords(inputs []int) int {
	inputs = getPasswordsWithAdjacentDigits(inputs)
	inputs = getPasswordWithoutDecreasingDigits(inputs)
	inputs = getPasswordsNoPartOfLargerGroups(inputs)

	return len(inputs)
}

func main() {
	inputs := generateInputs(147981, 691423)

	fmt.Printf("%d", getAmountOfMatchingPasswords(inputs))
}
