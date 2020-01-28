package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 1; i < 10; i++ {
		/*
		ニュートン法
		z^2-xはz^2がxからどれだけ離れているかを表す
		2zによる除算はz^2の導関数
		z^2の変化の速さ（接線の傾き）が大きいと、より大きくなる
		*/
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}