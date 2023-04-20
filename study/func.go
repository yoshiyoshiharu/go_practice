package main

import "fmt"

func main() {
  sum := add(2, 4)

	fmt.Println(sum)
	fmt.Println(addMinus(1, 2))

	// 複数返す関数のうち、不要な戻り値があるときは_を使う
	_, x := addMinus(1, 2)
	fmt.Println(x)
}

func add(x int, y int) int {
	return x + y
}

// 複数の値を返す関数
func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}
