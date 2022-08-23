package main

import "fmt"

func main() {
	firstName := "Jack"

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Daphne"
	fmt.Println(ptr, *ptr)
}
