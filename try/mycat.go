package main

// 作成するcatコマンドの仕様
// 引数でファイルパスの一覧を貰い、そのファイルを与えられた順に標準出力にそのまま出力するコマンドを作ってください。
// また、-nオプションを指定すると、行番号を各行につけて表示されるようにしてください。
// なお、行番号はすべてのファイルで通し番号にしてください。

import (
	"fmt"
	"flag"
	"os"
	"bufio"
)

func main() {
	var display_line_number *bool = flag.Bool("n", false, "displa line number flag")

	flag.Parse()
	fmt.Println(*display_line_number)
	file_paths := flag.Args()

	for i := 0; i < len(file_paths); i++ {
		printFileContent(file_paths[i], *display_line_number)
	}
}

func printFileContent(file_path string, display_line_number bool) error {
	file, err := os.Open(file_path)
	if err != nil { return err }
	defer file.Close()

	line := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if display_line_number {
		  fmt.Printf("%d: ", line)
		}

		fmt.Printf("%s\n", scanner.Text())
		line++
	}

	return err
}
