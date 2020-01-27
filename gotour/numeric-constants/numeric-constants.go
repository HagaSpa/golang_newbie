package main

import "fmt"

const (
	// 1を100桁左へシフト。1の後に0が100個続く
	Big = 1 << 100
	// Bigを99桁右へシフト。
	Small = Big >> 99
	// 1 << 3 は1から3桁左にずれるので100.これを2進数で表現すると8.
	Sample = 1 << 3
)

func needInt(x int) int { 
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	fmt.Println(Small)
	
	if Sample == 8 {
		fmt.Println("Sample is 8")
	}
}