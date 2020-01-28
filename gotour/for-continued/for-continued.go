package main

import "fmt"

func main() {
	sum := 1
	// 初期化式と後処理は任意
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}