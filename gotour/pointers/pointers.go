package main

import "fmt"

func main() {
	i, j := 42, 2701
	
	p := &i         // iのアドレスを取得し、ポインタ変数pへ格納
	fmt.Println(*p) // iのアドレス先の値を表示
	*p = 21         // iのアドレス先の値を21へ変更
	fmt.Println(i)  // iの値を表示

	p = &j          // jのアドレスを取得し、ポインタ変数pへ格納
	*p = *p / 37    // jのアドレス先の値を37で割る
	fmt.Println(j)  // jの値を表示
}