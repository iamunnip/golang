package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Println("Enter the first number")
	fmt.Scan(&a)
	fmt.Println("Enter the second number")
	fmt.Scan(&b)

	c = a - b

	fmt.Printf("Result = %d", c)
}