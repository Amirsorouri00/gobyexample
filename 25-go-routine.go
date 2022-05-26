package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	go f("go routine")

	time.Sleep(time.Second)
	// https://gobyexample.com/waitgroups
	// more robust way to wait for
	// go routine threads to finish
	fmt.Println("done")
}
