// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"

	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	fmt.Println(funcs.PopCount0(5356565))
	fmt.Println(funcs.PopCount1(5356565))
	fmt.Println(funcs.PopCount2(5356565))
	fmt.Println(funcs.PopCount3(5356565))
}
