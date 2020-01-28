package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// switch true と同義。if-then-elseの代わりになる
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}