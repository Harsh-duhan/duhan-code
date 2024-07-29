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

	fmt.Println("version 2")
	fmt.Println("version 3")
	fmt.Println("version 4")
	fmt.Println("version 5")
	fmt.Println("version 6")
	fmt.Println("version 7")
	fmt.Println("version 8")
	fmt.Println("version 10")
}
