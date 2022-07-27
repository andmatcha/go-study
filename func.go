package main

import (
	"fmt"
)

func saySth(str string) {
	fmt.Println(str)
}

func sum(x int, y int) int {
	return x + y
}

// func cal(x int, y int) (int, int) {
// 	a := x * y
// 	b := x / y

// 	return a, b
// }

func cal(x int, y int) (a int, b int) {
	a = x * y
	b = x / y

	return
}

func main() {
	saySth("Hello, world!")
	fmt.Println(sum(4, 6))

	result1, result2 := cal(10, 2)
	fmt.Println(result1, result2)

	// 無名関数
	func01 := func (sth string) {
		fmt.Println(sth)
	}
	func01("hello")

	func (sth string) {
		fmt.Println(sth)
	}("hello")
}
