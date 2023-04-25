package main

// String() string メソッドを持つインタフェースを作る
// そして3つ以上Stringerインタフェースを実装する型を作る
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

func Tostring(s Stringer) {
	println(s.String())
}

func main() {
	p := Person{
		name: "yoshida",
		age: 23,
	}

	Tostring(p)
}
