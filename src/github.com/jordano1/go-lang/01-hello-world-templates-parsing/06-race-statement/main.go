package main

import (
	"fmt"
	"runtime"
	"sync"
)

func foo() {

}

func main() {
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
