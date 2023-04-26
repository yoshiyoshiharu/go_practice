package main

import "fmt"
import "time"

// 並列処理 → 同じ処理を並行して実行 RubyのThread
// 並行処理 → 異なる処理を並行して実行 Goroutine

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
