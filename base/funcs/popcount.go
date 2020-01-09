package funcs

import (
	"fmt"
	"time"
)

var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

// PopCount0 returns the population count (number of set bits) of x.
func PopCount0(x uint64) int {
	start := time.Now()
	res := int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
	nsecs := time.Since(start).Nanoseconds()
	fmt.Printf("%d nsec ", nsecs)
	return res
}

// PopCount1 returns the population count (number of set bits) of x.
func PopCount1(x uint64) int {
	var i uint
	res := 0
	start := time.Now()
	for i = 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	nsecs := time.Since(start).Nanoseconds()
	fmt.Printf("%d nsec ", nsecs)
	return res
}

// PopCount2 returns the population count (number of set bits) of x.
func PopCount2(x uint64) int {
	res := 0
	start := time.Now()
	for x != 0 {
		res += int(x & 1)
		x = x >> 1
	}
	nsecs := time.Since(start).Nanoseconds()
	fmt.Printf("%d nsec ", nsecs)
	return res
}

// PopCount3 returns the population count (number of set bits) of x.
func PopCount3(x uint64) int {
	res := 0
	start := time.Now()
	for x != 0 {
		res++
		x = x & (x - 1)
	}
	nsecs := time.Since(start).Nanoseconds()
	fmt.Printf("%d nsec ", nsecs)
	return res
}
