package main

import(
	"fmt"
)

// 構造体を定義
type Student struct {
	name string // フィールドを定義
	math, english float64
}

func main() {
	// var struct01 Student
	// struct01.name = "jin"
	// struct01.math = 30
	// struct01.english = 80

	// struct01 := Student{"jin", 30, 80}

	struct01 := Student{name: "jin", math: 30, english: 80}

	fmt.Println(struct01)
}
