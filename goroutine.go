package main

import "fmt"
import "time"

func parallel_func() {
	for i := 0; i < 10; i++ {
		fmt.Println("parallel", i)
		time.Sleep(1 * time.Millisecond)
	}
}


func main() {
	go parallel_func()
	for i := 0; i < 10; i++ {
		fmt.Println("main", i)
		time.Sleep(1 * time.Millisecond)
	}
}
