// Server2 is a minimal "echo" and counter server.
package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/zyc737347123/goPractice/base/funcs"
)

func main() {
	fmt.Println(funcs.Comma1("-12345"))
	x1 := sha256.Sum256([]byte("ajjjjjj"))
	x2 := sha256.Sum256([]byte("ajjjjjjj"))
	fmt.Println(funcs.SHACount(&x1, &x2))
}
