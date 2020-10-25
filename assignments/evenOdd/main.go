package main

import "fmt"

func main() {
	var nums []int

	// create slice of ints from 0 to 10
	for i := 0; i <= 10; i++ {
		nums = append(nums, i)
	}

	// Iterate through the slice with a for loop.
	// For each element, check to see if the number is even or odd.
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}
