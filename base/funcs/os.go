package funcs

// Is64 returns true if this machince is 64 bit
func Is64() bool {
	if 32<<(^uint(0)>>63) == 64 {
		return true
	}
	return false
}
