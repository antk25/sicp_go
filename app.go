package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers = []int{2, 8, 1, 9, 11, 13, 18}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	greaterNumbers := []int{numbers[0], numbers[1]}

	fmt.Println(sumSquare(greaterNumbers))
}

func sumSquare(numbers []int) int {
	var sum int

	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i] * numbers[i]
	}

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}
