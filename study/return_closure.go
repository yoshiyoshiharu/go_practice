// クロージャを返す関数
package main

import "fmt"

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	twoBase := makeMult(2) // twoBaseは関数
	threeBase := makeMult(3)

	fmt.Println(twoBase(2))
	fmt.Println(twoBase(3))
	fmt.Println(threeBase(2))
	fmt.Println(threeBase(3))
}
