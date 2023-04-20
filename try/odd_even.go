package main

import "fmt"

// 奇数と偶数を判定
func main() {
	for num := 1; num <= 100; num++ {
		if num % 2 == 0 {
			fmt.Printf("%d-%s\n", num, "偶数")
		} else {
			fmt.Printf("%d-%s\n", num, "奇数")
		}
	}

	// switchで書く
	for num := 1; num <= 100; num++ {
		switch {
		case num % 2 == 0:
			fmt.Printf("%d-%s\n", num, "偶数")
		default:
			fmt.Printf("%d-%s\n", num, "奇数")
		}
	}
}
