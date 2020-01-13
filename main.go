// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"

	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	fmt.Println(funcs.Comma1("-12345"))
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println(funcs.Rotate(s, 2))
	strs := []string{"aaa", "aaa", "bbb", "bbb", "bb", "ccc"}
	fmt.Println(funcs.RemoveSame(strs))
}
