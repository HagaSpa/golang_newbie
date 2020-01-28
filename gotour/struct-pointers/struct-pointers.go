package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v   // 構造体のポインタ
	p.X = 1e9 // 構造体のフィールドXに、ポインタ経由でアクセスする. e9は0を9桁追加
	fmt.Println(v)
}