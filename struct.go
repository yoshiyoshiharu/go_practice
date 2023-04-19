package main

import "fmt"

func main() {
	var person Person
	person.SetPerson("太郎", 20)
	fmt.Println(person.GetPerson())

	// 初期化の書き方
	person2 := Person{"次郎", 30}
	person3 := Person{name: "三郎", age: 40}

	fmt.Println(person2.GetPerson())
	fmt.Println(person3.GetPerson())
}

// classの代わりに構造体を使う
type Person struct {
	name string
	age int
}

// インスタンスメソッド 関数名の前に(インスタンス変数名 *構造体名)
func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}
