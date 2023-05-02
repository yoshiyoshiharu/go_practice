package main

import "fmt"

type Hex int

// Hex型のメソッドを定義
// 関数定義の前にレシーバを渡す
func (h Hex) Tostring() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	var hex Hex = 100
	println(hex.Tostring())

	// メソッドは値としても扱える
	f := hex.Tostring
	println(f())

	// メソッド式
	f2 := Hex.Tostring
	println(f2(hex))
}
