package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string, 10)

	myMap["a"] = "b"
	value, exists := myMap["a"]
	if exists {
		println(value)
	}
	for k, v := range myMap {
		println(k, v)
	}
	fmt.Printf("myMap %+v\n\n", myMap)

	myFuncMap := map[string]func() int{
		"funcA": func() int { return 1 },
	}
	fmt.Println(myFuncMap)

	f := myFuncMap["funcA"]
	fmt.Println(f())
}
