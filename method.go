package main

import(
	"fmt"
)

// 構造体を定義
type Student struct {
	name string // フィールドを定義
	math, english float64
}

func (s Student) avg() {
	fmt.Println(s.name, (s.math + s.english) / 2)
}

func main() {
	a := Student{name: "jin", math: 30, english: 80}
	a.avg()
}
