package popcount

// go test -cpu=4 -bench=.

import (
	"testing"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountShift(x uint64) int {
	c := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			c++
		}
	}
	return c
}

func PopCountClear(x uint64) int {
	c := 0
	for x != 0 {
		x = x & (x - 1)
		c++
	}
	return c
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(0x1234567890ABCDEF)
	}
}
