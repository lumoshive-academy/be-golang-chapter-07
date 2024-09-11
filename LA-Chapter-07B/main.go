package main

import "fmt"

func main() {

	// for dengan kondisi tunggal
	count := 0
	for count < 5 {
		fmt.Println("Count:", count)
		count++
	}

	// for klausa loop klasik
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// for invinite loop
	for {
		fmt.Println("Count:", count)
		count++
		if count == 5 {
			break
		}
	}

	// for with range
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// break statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("i:", i)
	}

	// continue statement
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("i:", i)
	}

}
