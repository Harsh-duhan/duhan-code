package main

import "fmt"

func main() {
	fmt.Println("pointers")
	var roll1 int = 1
	var q *int = &roll1

	var roll2 int = 2
	var P *int = &roll2
	var roll3 int = 3
	var ptr *int = &roll1

	fmt.Println("value of roll", roll1)
	fmt.Println("value of pointer ", q)
	fmt.Println("value of roll", roll2)
	fmt.Println("value of pointer ", P)
	fmt.Println("value of roll", roll3)
	fmt.Println("value of pointer ", ptr)
}
