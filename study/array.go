package main

import "fmt"

func main() {
	// 配列 あらかじめ個数を定義し、あとから変更できない
	var array1 [3]int           // 容量が3の配列
	array2 := [3]int{1, 2, 3}   // 初期値
	array3 := [...]int{1, 2, 3} // 容量を省略すると自動で容量を決定
	println(array1[0], array2[0], array3[0])

	// スライス 個数が変更可能
	// 配列の一部を切り出したデータ構造で背後に配列が存在する
	sl1 := []int{1, 2, 3}
	sl2 := make([]int, 3, 10) // 長さと容量を指定して作成
	println(sl1, sl2)         // => [3/3]0x14000096e50 [3/10]0x14000096e68

	// スライスの操作
	println(len(sl1), cap(sl1)) // 長さと容量を取得
	sl1 = append(sl1, 4, 5)     // 追加
	println(len(sl1), cap(sl1))

	// マップ:連想配列 map[キーの型]値の型
	map1 := map[string]int{
		"key":  200,
		"key2": 300, // カンマ必須
	}
	map2 := make(map[string]int)
	map3 := make(map[string]int, 10) // 容量を指定して作成
	map4 := map[string]int{"x": 100, "y": 200}

	map1["key3"] = 400  // 挿入
	delete(map1, "key") // 削除

	println(map1, map2, map3, map4)

	// mapのループ
	for key, value := range map1 {
		fmt.Printf("%s = %d \n", key, value)
	}
}
