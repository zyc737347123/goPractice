// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"

	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	fmt.Println(funcs.Comma1("-12345"))
	fmt.Println(funcs.Comma2("12345.3242"))
	fmt.Println(funcs.CompareByte("3322", "2332"))
}
