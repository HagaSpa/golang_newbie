package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// func内部では:=を使用してvar宣言を省略できる
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}