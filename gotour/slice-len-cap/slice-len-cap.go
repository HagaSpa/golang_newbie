package main

import "fmt"

/*
sliceの長さは要素数。
sliceの容量は、sliceの元となる配列の要素数。

要素数を超えた参照や、容量を超えた拡張などはruntime errorが起きる
*/
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 要素数0のsliceを作成
	s = s[:0]
	printSlice(s)

	// 要素を拡張。最初に定義した6以上は参照できない。
	s = s[:4]
	printSlice(s)

	// 最初から2つの要素を削除
	s = s[2:]
	printSlice(s)

	// runtime error: slice bounds out of range [:5] with capacity 4
	// s = s[:5]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}