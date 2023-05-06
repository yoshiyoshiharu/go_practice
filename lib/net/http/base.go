package main

import(
	"fmt"
	"net/http"
)

// localhost:8080にアクセスすると文字列が表示される
func main() {
	http.HandleFunc("/", handler) // ハンドラを登録(パターン, 関数)
	http.ListenAndServe(":8080", nil) // HTTPサーバーの起動(ホスト名:ポート番号, HTTPハンドラ(nilのときは登録したハンドラ))
}

// httpハンドラ
// 第1引数はレスポンスを書き込む先
// 第2引数はクライアントからのリクエスト
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello") // 書き込み先を指定して出力
	fmt.Fprintln(w, r.FormValue("msg")) // FormValueでリクエストパラメタを取得
}
