package main

import (
	"fmt"

	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	f, f2 := funcs.Fibonacci(), funcs.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(), f2())
	}
}
