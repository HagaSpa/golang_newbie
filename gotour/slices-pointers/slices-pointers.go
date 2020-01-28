package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	// スライスは配列の一部を参照している、ポインタなので、スライスの要素を変更すると元の配列に影響がある
	// 元の配列に影響があるということは、そこを参照する別のスライスも影響される
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}