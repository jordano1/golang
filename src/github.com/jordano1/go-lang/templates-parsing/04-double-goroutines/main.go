package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("hey")
	fmt.Println("hi")
	fmt.Println("goroutines:", runtime.NumGoroutine())
	fmt.Println("ho")
	wg.Add(1)
	go foo()
	fmt.Println("goroutines:", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
