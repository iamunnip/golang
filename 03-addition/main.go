package main

import "fmt"

func main() {

	// declaring variables
	var a int
	var b int
	var c int

	// getting input from user
	fmt.Println("Enter the first number")
	fmt.Scanf("%d \n",&a)
	fmt.Println("Enter the second number")
	fmt.Scanf("%d \n",&b)

	// adding the numbers
	c = a + b

	// displaying the result
	fmt.Printf("Result = %d", c)
}

/*
$ go run main.go
Enter the first number
1
Enter the second number
2
Result = 3
*/