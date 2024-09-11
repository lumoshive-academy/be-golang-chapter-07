package main

import "fmt"

func main() {
	age := 18

	// basic if expression
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	// if else expression
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult.")
	}

	// if with initializer
	if age := 18; age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult.")
	}

	// switch statement
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Today is Monday.")
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	default:
		fmt.Println("Today is not Monday or Tuesday.")
	}

	// switch statement multiple cases
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend.")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	default:
		fmt.Println("Unknown day.")
	}

	// switch statement without expression
	switch {
	case age < 13:
		fmt.Println("You are a child.")
	case age < 20:
		fmt.Println("You are a teenager.")
	default:
		fmt.Println("You are an adult.")
	}

	fmt.Println("End of program.")
}
