package main

import (
	"errors"
	"fmt"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]

	return name, ok
}

// SimpleDataStoreのファクトリ関数
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

// ロジックがDataStoreとLoggerに依存することを表明
// LogOutputやSimpleDataStroeに依存したくない → あとから入れ替えられる
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

// このままではLogOutputがLoggerとして使えないのでアダプターとなる関数型を定義
type LogAddapter func(message string)

func (lg LogAddapter) Log(message string) {
	lg(message)
}

// 依存性は定義されたので、ビジネスロジックの定義
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("SayHello(" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)

	if !ok {
		return "", errors.New("ユーザーが見つかりませんでした。")
	}
	return name + "こんにちは。", nil
}

// SimpleLogicのファクトリ関数
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

func main() {
	l := LogAddapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)

	message, _ := logic.SayHello("1")
	println(message)
	message, _ = logic.SayHello("2")
	println(message)
	message, _ = logic.SayHello("3")
	println(message)
}
