package main

func makeChan() chan int {
	return make(chan int)
}

// 双方向チャネル
func twowayChan(ch chan int) int {
	// ch <- 200 これはエラー: チャネルの送信はゴルーチン内でやる
	go func() { ch <- 200 }()
	return <- ch
}

// <- chan int は受信専用チャネル
func recieveChan(recieve <- chan int) int {
	return <- recieve
}

func main() {
	ch1 := makeChan()
	println(twowayChan(ch1))

	ch2 := makeChan()
	// chan <- int は送信専用チャネル
	go func(ch chan <- int) { ch <- 100 }(ch2)
	println(recieveChan(ch2))
}
