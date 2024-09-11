package main

import "fmt"

// basic function
func sayHello() {
	fmt.Println("Hello, World!")
}

// function with parameter
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// function with return
func add(a int, b int) int {
	return a + b
}

// function with multiple return
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	// call Function
	sayHello()

	// call function paramter
	greet("Alice")
	greet("Bob")

	// call function with return
	result := add(3, 4)
	fmt.Println("Result:", result)

	// call function with multiple return
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q)
	fmt.Println("Remainder:", r)
}
