package funcs

// Fibonacci is a func
func Fibonacci() func() int {
	a0, a1 := 0, 1
	return func() int {
		res := a1
		a0, a1 = a1, a0+a1
		return res
	}
}
