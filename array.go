package main

import "fmt"

func main() {
	// 配列 あらかじめ個数を定義し、あとから変更できない
	var a1 [3]int
	fmt.Println(a1)

	a2 := [3]int{1, 2, 3}
	fmt.Println(a2)

	// スライス 個数が変更可能
  a3 := []int{1, 2, 3}
	fmt.Println(a3)

	fmt.Println(len(a3), cap(a3))  // lenは実際に使用されているメモリ、capはメモリ上に確保されているメモリ

	// マップ:連想配列 map[キーの型]値の型

	a4 := map[string]int{
		"key": 200,
		"key2": 300, // カンマ必須
	}

	fmt.Println(a4)

	a4["key3"] = 400 // 挿入
	delete(a4, "key") // 削除

	fmt.Println(a4)

	// mapのループ
	for key, value := range a4 {
		fmt.Printf("%s = %d \n", key, value)
	}
}
