package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	// 変数の定義
	var a int // 型の定義

	var c int = 1 // 型の定義と同時に値を代入
	var d = 1     // 引数が明示的なときは型の定義を省略
	e := 1        // 引数があるときはvarを省略

	// 変数をまとめて定義
	var (
		f int = 1
		g int = 2
	)

	// 複数の変数を同時に定義
	var name, age = "Taro", 20

	fmt.Println("vars:", a, c, d, e, f, g, name, age)

	// 定数の定義
	const FOO = 100

	// 型変換
	var num int = 123
	var big_num = uint64(num)
	fmt.Println(num, big_num)

	// 文字列
	var char rune = 'あ'         // 1文字
	var string string = "あいうえお" // 1文字以上

	fmt.Println(char, string)
}
