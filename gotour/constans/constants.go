package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Printf("World Type is %T\n", World)

	fmt.Println("Happy", Pi, "Day")
	fmt.Printf("Pi Type is %T\n", Pi)  	// constの型は推論される。明示的に型の宣言はできない。

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Printf("Truth Type is %T\n", Truth)
}