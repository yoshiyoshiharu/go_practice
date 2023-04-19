package main

import "fmt"

// Interface: ポリモーフィズム

type Printable interface {
	Tostring() string // Tostring()を呼び出すとstringを返す
}

// Printoutを共通化
func Printout(p Printable) { // Printable型の引数
	fmt.Println(p.Tostring()) // pはTostring()に応答する
}

type Person struct {
	name string
	age int
}

type Book struct {
	title string
	price int
}

// PrintableになるにはTostring()を実装する必要がある
// p *Personの米印の有無でどう違うのか
func (p Person) Tostring() string {
	return p.name
}

func (b Book) Tostring() string {
	return b.title
}

func main() {
	person := Person{name: "Taro", age: 20}
	book := Book{title: "Go言語", price: 1000}

	Printout(person)
	Printout(book)
}
