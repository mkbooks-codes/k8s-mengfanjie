package main

import "fmt"

func main() {
	// 基本形式
	condition1 := false
	condition2 := true
	if condition1 {
		// do something
		fmt.Printf("condition1 %t", condition1)
	} else if condition2 {
		// do something else
		fmt.Printf("condition2 %t", condition2)
	} else {
		// catch-all or default
		fmt.Printf("else")
	}

	// if 的简短语句
	fmt.Println()
	x := 99
	if v := x - 100; v < 0 {
		fmt.Printf("v: %d", v)
	}

	fmt.Println()

	// switch
	var1 := 2
	val1 := 1
	val2 := 2
	val3 := 3
	switch var1 {
	case val1: //空分支
		fmt.Println("val1")
	case val2:
		fallthrough //执行case3中的f()
	case val3:
		f()
	default: //默认分支
		fmt.Println("default")
	}
}

func f() {
	fmt.Println("val3")
}
