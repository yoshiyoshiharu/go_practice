package main

import "fmt"

type Printable interface {
	Tostring() string // Tostring()を呼び出すとstringを返す
}

type Person struct {
	name string
}

func (p Person) Tostring() string {
  return p.name
}

// interface{} 任意の型の関数を受け取る
func PlusOne(x interface{}) int {
	// 変数.(型名)でany型の変数を型キャストする
	return x.(int) + 1
}

func Printout(x interface{}) {
	q, ok := x.(Printable)

	if ok {
		fmt.Println(q.Tostring())
	} else {
		fmt.Println("Not Printable")
	}
}
func main() {
	fmt.Println(PlusOne(1))

	person := Person{name: "Taro"}
	Printout(person)
	Printout(1)
}
