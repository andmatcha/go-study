package main // プログラムは何かしらのパッケージに属している必要がある。mainパッケージは必ずひとつは必要

import ( // 使用するパッケージをインポートする。未使用だとコンパイルエラー。
	"fmt"
	"reflect"
)

func main() { // mainパッケージのmain関数がエントリポイント
	num01 :=1 // 変数の宣言と初期化は:=でできる。変数も未使用だとコンパイルエラー。
	var num02 int = 12345 // 型を明示的に宣言することもできる。

	string01 := "こんにちは" // 文字列型も同様
	var string02 string = "こんばんは"

	var bool01 bool = string01 == string02 // bool型

	arr01 := [3]string{"apple", "orange", "peach"} // 配列型 [要素数]データ型{要素, 要素, ...}
	arr02 := [...]string{"apple", "orange", "peach"} // 要素数は[...]で省略できる
	arr03 := [2][2]int{{ 1, 2 }, {3, 4}} // 多次元配列

	fmt.Println("################################")
	fmt.Println(reflect.TypeOf(num01)) // fmt.Println()を使って標準出力できる。
	fmt.Println(reflect.TypeOf(num02)) // reflect.TypeOf()は変数の型を返す。
	fmt.Println("################################")
	fmt.Println(string01)
	fmt.Println(string02)
	fmt.Println(bool01)
	fmt.Println("################################")
	fmt.Println(arr01[0])
	fmt.Println(arr01[1])
	fmt.Println(arr01[2])
	fmt.Println(arr02)
	fmt.Println(arr03[1][0])
	fmt.Println("################################")


	x := 2
	y := 5
	fmt.Println(x + y) // 7
	fmt.Println(x - y) // -3
	fmt.Println(x * y) // 10
	fmt.Println(x / y) // 0
	fmt.Println(x % y) // 2
	fmt.Println(x > y) // false
	fmt.Println(x <= y) // true
	fmt.Println(x == 2 && y != 3) // true
	fmt.Println(x == 5 || y == 5) // true
	fmt.Println("################################")

	x += 2
	fmt.Println(x) // 4
	x++
	fmt.Println(x) // 5
	x--
	fmt.Println(x) // 4
	fmt.Println("################################")

	// 条件分岐
	age01 := 18

	if age01 >= 20 {
		fmt.Println("adult")
	} else if age01 == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}


	if age02 := 20 + 1; age02 >= 20 { // 簡易的な文も書ける。
		fmt.Println("adult")
	} else if age02 == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}

	// 繰り返し処理
	for i := 1; i <= 10; i++ {
		if i % 3 == 0 {
			fmt.Println("fizz")
			continue
		} else if i == 8 {
			break
		}
		fmt.Print(i, "\n")
	}

}
