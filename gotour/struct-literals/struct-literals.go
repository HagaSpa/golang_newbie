package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // Xを1、Yを2で初期化
	v2 = Vertex{X: 1}  // Xだけを指定。Yは0で初期化
	v3 = Vertex{}      // XもYも0
	p  = &Vertex{1, 2} // *Vertexのポインタ。v1,v2,v3どれとも違う別の領域
)

func main() {
	p.X = 2
	fmt.Println(v1, p, v2, v3)
}
