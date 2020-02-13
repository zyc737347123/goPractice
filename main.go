package main

import (
	"fmt"

	"github.com/zyc737347123/goPractice/base/interfaces"
)

func main() {
	wo := new(interfaces.WordCounter)
	w, v := interfaces.CountingWriter(wo)
	w.Write([]byte("hello hi fdjlkf\nlink"))
	fmt.Println(*v)
}
