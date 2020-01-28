package main

import "fmt"

func main() {
	fmt.Println("counting")

	/*
	複数deferを行うとそれはスタックに積み上げられる。LIFO
	*/
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}