package funcs

// Fibonacci is a func
func Fibonacci() func() int {
	a0 := 0
	a1 := 1
	n := 0
	return func() int {
		n = n + 1
		if n-1 == 0 {
			return a0
		}
		if n-1 == 1 {
			return a1
		}
		a2 := a0 + a1
		a0 = a1
		a1 = a2
		return a2
	}
}