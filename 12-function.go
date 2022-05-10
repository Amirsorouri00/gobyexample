package main

import "fmt"

func main() {

	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}

func plusPlus(i1, i2, i3 int) int {
	return plus(i3, plus(i1, i2))
}

func plus(i1, i2 int) int {
	return i1 + i2
}
