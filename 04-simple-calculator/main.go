package main

import "fmt"

func addition(a, b int) int {
	return a + b
}

func substraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) int {
	return a / b
}

func main() {

	var x, y, z int

	fmt.Println("Enter the first number")
	fmt.Scanf("%d\n", &x)
	fmt.Println("Enter the second number")
	fmt.Scanf("%d\n", &y)

	fmt.Println("Menu")
	fmt.Println("1. Addition")
	fmt.Println("2. Substraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	fmt.Println("Enter your choice")
	fmt.Scanf("%d\n", &z)

	switch z {

	case 1:
		fmt.Println("Result =", addition(x, y))

	case 2:
		fmt.Println("Result =", substraction(x, y))

	case 3:
		fmt.Println("Result =", multiplication(x, y))

	case 4:
		fmt.Println("Result =", division(x, y))

	default:
		fmt.Println("Error")
    }
}