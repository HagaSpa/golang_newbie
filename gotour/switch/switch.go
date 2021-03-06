package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "lnux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}