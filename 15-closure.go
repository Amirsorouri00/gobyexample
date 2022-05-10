package main

import "fmt"

// Go supports anonymous functions (
// 		https://en.wikipedia.org/wiki/Anonymous_function
// ), which can form closures.
// Anonymous functions are useful when you want to define a function inline without having to name it.

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSeq()
	fmt.Println(newInt())
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
