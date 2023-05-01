package main

// 型キャスト
func main() {
	sum = 5 + 6 + 3
	var avg float64
	avg = float64(sum) / float64(3) // float64で統一する必要がある。

	if avg > 4.5 {
		println("%f", avg)
	}
}
