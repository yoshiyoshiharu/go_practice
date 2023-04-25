package main
// Stringerインタフェースを引数で受け取る関数を作る
// 受け取った値を上記の3つの具象型によって分岐し、具象型の型名と値を表示する

type Stringer interface {
	String() string
}

type Car struct {
	name string
	price int
}

type Book struct {
	title string
	price int
}

type Person struct {
	name string
	age int
	sex int
}

func (c Car) String() string {
	return c.name
}

func (b Book) String() string {
	return b.title
}

func (p Person) String() string {
	return p.name
}

func DisplayStringer(s Stringer) {
	switch s.(type) {
	case Car:
		println("I am Car:")
	case Book:
		println("I am Book")
	case Person:
		println("I am Person")
	default:
		println("I am Other:")
	}
	println(s.String())
}

func main() {
	p := Person{
		name: "yoshida",
		age: 23,
	}
	b := Book{
		title: "book",
		price: 120,
	}
	c := Car{
		name: "car kind",
	}

	DisplayStringer(p)
	DisplayStringer(b)
	DisplayStringer(c)
}
