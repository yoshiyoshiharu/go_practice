// 可変長引数

package main

import "fmt"

// valsは可変長引数でスライスとして扱える
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals)) // 長さ0で容量がvalsの長さのスライス
	for _, v := range vals {
		out = append(out, v + base) // appendは同じ型のスライスを返す
	}
	return out
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(addTo(3, a...)) // 4, 5, 6, 7
}
