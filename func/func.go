package main

import "fmt"

func main() {
	sum := add(2, 4)

	fmt.Println(sum)
	fmt.Println(addMinus(1, 2))

	// 複数返す関数のうち、不要な戻り値があるときは_を使う
	_, x := addMinus(1, 2)
	fmt.Println(x)

	// 関数型
	fs := make([]func() string, 2) // 関数を要素に持つスライス
	fs[0] = func() string { return "hello" }
	fs[1] = func() string { return "world" }

	for _, f := range fs {
		println(f())
	}

	deferHello()
}

func add(x int, y int) int {
	return x + y
}

// 複数の値を返す関数
func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

// deferは関数終了時に実行(returnの後)
// スタック形式LIFOで実行される(最後に指定したものが最初に実行)
func deferHello() {
	msg := "defer1"
	defer println(msg)
	msg = "defer2"
	defer println(msg)
	println("defer3")
}
