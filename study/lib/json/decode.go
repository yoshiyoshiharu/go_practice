package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	var p Person

	// デコード → JSONからGo構造体
	// json.Unmarshalの利用(第一の選択肢)
	jsonString := `{"name": "yoshida", "age": 23}`

	if err := json.Unmarshal([]byte(jsonString), &p); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", p)

	// io.Readerの形で得ているのならばデコーダを用意する方法がある
	dec := json.NewDecoder(strings.NewReader(jsonString))

	var p2 Person

	if err := dec.Decode(&p2); err != nil {
		println(err)
		return
	}

	fmt.Printf("%+v\n", p2)
}
