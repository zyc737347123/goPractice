package main

import (
	"fmt"

	"github.com/zyc737347123/goPractice/base/funcs"
	"github.com/zyc737347123/goPractice/base/structs"
)

func main() {
	var test *structs.IntSet = new(structs.IntSet)
	test.Add(1)
	test.Add(3)
	test.Add(5)
	fmt.Println(test, test.Len())
	test.Remove(3)
	fmt.Println(test, test.Len())
	test.Add(4)
	fmt.Println(test, test.Len())
	t2 := test.Copy()
	fmt.Println(t2, t2.Len())
	test.Clear()
	fmt.Println(test, test.Len())
	fmt.Println(t2, t2.Len())
	test.UnionWith(t2)
	fmt.Println(test, test.Len())
	fmt.Println(test.Elems())

	// to know 64 or 32 bit machice
	fmt.Println(funcs.Is64())
}
