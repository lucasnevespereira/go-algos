package main

import "fmt"

func main() {
	var numbers []int = []int{5, 4, 2, 3, 1, 0}

	fmt.Println("Initial numbers:", numbers)

	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)

}

// The Sort
func bubbleSort(numbers []int) {

	var N int = len(numbers)

	var i int
	for i = 0; i < N; i++ {
		swap(numbers)
	}
}

// The Swap
func swap(numbers []int) {

	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	// while secondIndex is smaller than numbers len
	for secondIndex < N {

		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber {
			// swap the numbers
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		// Increase indexes
		firstIndex++
		secondIndex++
	}
}
