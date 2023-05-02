package main

import "fmt"
import "time"

// チャネルを用いてゴルーチンの終了を待ち合わせ
// 複数のごルーチンから共通の変数にアクセスするとスレッドセーフではない
// 並行処理の順序を制御するためにチャネル
func parallel_func(channelA chan<- string) {
	time.Sleep(3 * time.Second)
	channelA <- "3 seconds past!!!!" // <-はチャネルの送受信
}

func main() {
	channelA := make(chan string) // チャンネルを作成(stirngを返す)
	defer close(channelA)         // 終わったらチャンネルを閉じる
	go parallel_func(channelA)    // チャネルをゴルーチンに渡す
	message := <-channelA         // チャネルから値を受け取る
	fmt.Println(message)
}
