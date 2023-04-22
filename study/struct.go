package main

import "fmt"

// ユーザー定義型
type MyInt int

type Person struct {
	name string
	age int
	Birthday string
}

// クラスメソッドに相当する関数は関数名の前に (thisに相当する変数 *構造体名)
func (p *Person) setPerson(name string, age int, birthday string) {
	p.name = name
	p.age = age
	p.Birthday = birthday
}

func (p *Person) getPerson() (string, int, string) {
	return p.name, p.age, p.Birthday
}

func main() {
  var p1 Person

	p1.setPerson("Taro", 20, "1990/01/01")

	fmt.Println(p1.name, p1.age, p1.Birthday)
}
