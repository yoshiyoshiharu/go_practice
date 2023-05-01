package main

import "fmt"

func main() {
	var a []int // スライスのゼロ値はnil
  println(a) // [0/0]0x0
	println(a == nil) // true

	// append
	a = append(a, 10)
	a = append(a, 1, 2, 3, 4) // 複数の値を挿入

	var b = []int{5, 6, 7, 8}
	a = append(a, b...) // スライスを展開して挿入

	// make(type, len, cap)であらかじめ容量を指定する
	var c = make([]int, 3, 10)
	fmt.Println(c) // [0, 0, 0]

	// スライスのスライス
	// サブスライスはスライスとメモリを共有している
	d := []int{1, 2, 3, 4}
	fmt.Println(d[1:3]) // [2, 3]
	fmt.Println(d[:3])// [1, 2, 3]
	fmt.Println(d[1:]) // [2, 3, 4]
}
