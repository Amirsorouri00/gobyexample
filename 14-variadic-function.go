package main

import "fmt"

// Variadic functions (https://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.
func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
