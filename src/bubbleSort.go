package main

import (
	"fmt"
	"os"
	"strconv"
)

const NUM_OF_INPUTS = 10

func PrintSlice(numbers []int) {

	for i := 0; i < NUM_OF_INPUTS; i++ {
		fmt.Printf("%d, ", numbers[i])
	}
}

func BubbleSort(numbers []int) {
	swapped := true
	// Check when to stop
	for swapped {
		swapped = false
		// We will start from 1 since it will not be out of bounds
		for i := 1; i < NUM_OF_INPUTS; i++ {
			fmt.Println(numbers[i])
			if numbers[i-1] > numbers[i] {
				Swap(numbers, i)
				swapped = true
			}
		}
	}
}

func Swap(numbers []int, index int) {
	numbers[index], numbers[index-1] = numbers[index-1], numbers[index]
}

func main() {
	numbersSlice := make([]int, 0)
	var input string

	for i := 0; i < NUM_OF_INPUTS; i++ {
		fmt.Println("Enter a number: ")
		fmt.Scan(&input)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Bad Input!")
			os.Exit(1)
		}

		numbersSlice = append(numbersSlice, num)
	}

	BubbleSort(numbersSlice)
	PrintSlice(numbersSlice)
}
