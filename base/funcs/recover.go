package funcs

import "fmt"

// Recover is test func, will return non-nil but not use `return`
func Recover() (err error) {
	defer func() {
		p := recover()
		err = fmt.Errorf("res: %v", p)
	}()
	panic(1)
}
