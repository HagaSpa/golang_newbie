package main

import "fmt"

func main() {
	// array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// sはslice. primesの要素を1~3まで取得して作成
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Printf("%T\n", primes)
	fmt.Printf("%T\n", s)
}