package main

import(
  "fmt"
	"os"
)

type Stringer interface{
	String() string
}

type Person struct {
	name string
  age int
}

func (p Person) String() string {
	return p.name
}

type MyError string

// エラー型を定義するにはError()メソッドを実装する必要ある
func (e MyError) Error() string {
	return string(e)
}

func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, MyError("My Error occured")
}

func main() {
	p := Person{
		name: "haruki",
		age: 12,
	}

	if s, err := ToStringer(p); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR", err)
	} else {
		println("OK!", s.String())
	}
}
