package main

import "fmt"

func main() {
	var a1 int = 123
	var p1 *int // ポインタ変数

	p1 = &a1         // ポインタ変数にアドレスを代入
	fmt.Println(p1)  // アドレスが表示される
	fmt.Println(*p1) // 123 デリファレンス

	var a2 int = 456
	var a3 int = 456

	override(a2, &a3) // a2は値渡し、a3は参照渡し
	fmt.Println(a2, a3)
}

func override(a int, b *int) {
	a = 789
	*b = 789
}
