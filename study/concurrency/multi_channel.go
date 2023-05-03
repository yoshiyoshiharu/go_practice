// ゴルーチンの送受信を多重化
package main

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() { ch1 <- 100 }()  // 送信
	go func() { ch2 <- "hi" }() // 送信

	// 上から順番に評価されない
	// 複数のcaseが実行可能なときは、どれか一つがランダムに実行される
	select {
	case v1 := <-ch1: // 受信
		println(v1)
	case v2 := <-ch2: // 受信
		println(v2)
	}
}
