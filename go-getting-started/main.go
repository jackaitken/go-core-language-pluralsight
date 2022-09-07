package main

import (
	"fmt"

	"github.com/jackaitken/go-core-language-pluralsight/go-getting-started/models/go-getting-started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Arlo",
		LastName:  "Aitken",
	}

	fmt.Println(u)
}
