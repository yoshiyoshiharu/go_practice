package adder

import (
	"testing"
	"fmt"
	"time"
	"os"
)

var currentTime time.Time // パッケージブロックで宣言

func TestMain(m *testing.M) {
	fmt.Println("prepare test")
	currentTime = time.Now()
	exitVal := m.Run() // テストの実行
	fmt.Println("cleanup test")
	os.Exit(exitVal)
}

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("結果が違う: 期待する結果 5, 得られた結果", result)
	}

	fmt.Println(currentTime)
}
