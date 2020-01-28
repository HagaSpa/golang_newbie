package main

import "fmt"

func main() {
	/*
	deferは呼び出し元(main)が、終わる(return)まで実行されない.
	ただしdeferへ渡した引数の評価は過ぎに行われる
	*/
	defer fmt.Println("world")

	fmt.Println("hello")
}