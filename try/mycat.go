package main

// 作成するcatコマンドの仕様
// 引数でファイルパスの一覧を貰い、そのファイルを与えられた順に標準出力にそのまま出力するコマンドを作ってください。
// また、-nオプションを指定すると、行番号を各行につけて表示されるようにしてください。
// なお、行番号はすべてのファイルで通し番号にしてください。

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	var display_line_number *bool = flag.Bool("n", false, "displa line number flag")

	flag.Parse()
	fmt.Println(*display_line_number)
	file_paths := flag.Args()
	fmt.Println(file_paths)
}

func printFileContent(file_path string) error {
	file, err := os.Open(file_path)
	if err != nil { return err }
	defer file.Close()

	data := make([]byte, 1024) // バイト型スライス
	for {
		byte_size, err := file.Read(data)

		if byte_size == 0 {
			break
		}
		if err != nil {
			break
		}

		fmt.Println(string(data[:byte_size])) // バイト型スライスを文字列に変換して表示
	}
	return err
}
