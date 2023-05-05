package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	p := Person{
		Name: "Yoshida",
		Age: 23,
	}
	// エンコーディング → Go構造体からJSON
	// json.Marshal 返り値はエンコード結果がbytep[]で格納される
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)

	// Encoderの利用
	err2 := json.NewEncoder(os.Stdout).Encode(p)
	if err2 != nil {
		fmt.Println(err2)
	}
}
