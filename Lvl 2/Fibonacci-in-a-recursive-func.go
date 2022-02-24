package main

import "fmt"

func main() {
	var inputNumber int
	fmt.Scanln(&inputNumber)
	returnedFibonacci := fibonacci(inputNumber)
	fmt.Printf("You entered: %v",inputNumber)
    fmt.Println()
    fmt.Println("----------------------------")
    fmt.Printf("The result is: %v",returnedFibonacci)
}

func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
