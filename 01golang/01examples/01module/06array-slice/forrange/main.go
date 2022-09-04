package main

import (
	"fmt"
)

func main() {
	mySlice := []int{10, 20, 30, 40, 50}

	for _, value := range mySlice {
		value *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)

	// go vet
	// for index, _ := range mySlice {
	for index := range mySlice {
		mySlice[index] *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)
}
