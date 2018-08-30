package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("runtime:%s\narchitecture:%s", runtime.GOOS, runtime.GOARCH)
}
