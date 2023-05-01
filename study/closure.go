// 関数は値
// 関数内で定義された関数をクロージャという

package main

import(
	"fmt"
	"sort"
)

type Person struct{
	FirstName string
	LastName string
	Age int
}

func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fried", "Fredson", 18},
	}

	// 第2引数がクロージャ(関数引数)
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println("LastNameでソート", people)
}
