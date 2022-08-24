package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	slice2 := slice[1:]
	slice3 := slice[:2]
	slice4 := slice[1:2]

	fmt.Println(slice, slice2, slice3, slice4)
}
