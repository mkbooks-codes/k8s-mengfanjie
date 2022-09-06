package main

import "fmt"

func main() {
	input := "a"
	err, result := DuplicateString(input)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	err2, result2 := DuplicateString2(input)
	if err2 == nil {
		fmt.Println(result2)
	} else {
		fmt.Println(err2)
	}
}

// DuplicateString 与 DuplicateString2 等价。返回参数可以在返回值列表里提前定义。
func DuplicateString(input string) (err error, result string) {
	if input == "a" {
		err = fmt.Errorf("a is not allowed")
		return
	}
	result = input + input
	return
}

func DuplicateString2(input string) (error, string) {
	if input == "a" {
		return fmt.Errorf("a is not allowed"), ""
	}
	return nil, input + input
}
