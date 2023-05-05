package main

import(
	"net/http"
)

func main() {
	personMux := http.NewServeMux()
	personMux.HandleFunc("/greet",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("こんにちは"))
		})

	dogMux := http.NewServeMux()
	dogMux.HandleFunc("/greet",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("かわいい子犬ちゃん"))
		})

	mux := http.NewServeMux()
	// /personのパターンはpersonMuxが処理
	mux.Handle("/person/", http.StripPrefix("/person", personMux))
	// /dogのパターンはdogMuxが処理
	mux.Handle("/dog/", http.StripPrefix("/dog", dogMux))

	http.ListenAndServe(":8080", nil)
}
