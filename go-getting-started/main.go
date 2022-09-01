package main

import "fmt"

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Jack",
		LastName:  "Aitken",
	}

	fmt.Println(u)
}
