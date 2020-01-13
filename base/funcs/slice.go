package funcs

// Reverse reverses a slice of ints in place.
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Rotate is rotate to left
func Rotate(s []int, r int) []int {
	lens := len(s)
	//创建一个空的指定长度的slice
	res := make([]int, lens)
	for i := 0; i < lens; i++ {
		index := i + r
		if index >= lens {
			index = index - lens
		}
		res[i] = s[index]
	}
	return res
}

// RemoveSame ia a func
func RemoveSame(s []string) []string {
	length := len(s)
	index := 0
	var newIdx int

	for index < length {
		if (index+1 < length) && (s[index] == s[index+1]) {
			index++
		} else {
			newIdx++
			index++
			if (index < length) && (newIdx != index) {
				s[newIdx] = s[index]
			}
		}
	}

	return s[:newIdx]

}
