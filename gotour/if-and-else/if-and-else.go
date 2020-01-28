package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	// printlnで表示されるのは、両方のpow()の処理が終わってからのように見える
	// 23行目はelseでprintfしているが、それが最初に表示されるので、そう考えるのが自然
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}