package main

import (
	"fmt"
	"sync"
)

/*
go routineの中で複数のメソッドを呼び出してみる
*/

func createStr(i int) string {
	s := fmt.Sprintf("message: %d", i)
	return s
}

func printStr(s string) {
	fmt.Println(s)
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	wg := &sync.WaitGroup{}

	for _, v := range arr {
		wg.Add(1)
		// この処理がgoroutineでそれぞれ実行される。ここでは10並列
		go func(i int)  {
			s := createStr(i)
			printStr(s)
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println("Done...")
}