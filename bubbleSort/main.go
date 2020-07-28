package main

import "fmt"

func main() {
	var numbers []int = []int{5, 4, 2, 3, 1, 0}
	fmt.Println("Our list of numbers is:", numbers)
	swap(numbers)
}

func swap(numbers []int) {

	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	// while secondIndex is smaller than numbers len
	for secondIndex < N {

		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		fmt.Println("Comparing : ", firstNumber, secondNumber)

		// Increase indexes
		firstIndex++
		secondIndex++
	}
}
