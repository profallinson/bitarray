package bitarray

import (
	"fmt"
)

func PrintMatrix(deg float64, s string, size uint64) {
	start := uint64(0)
	end := size
	fmt.Println(deg)
	for end <= size*size {
		fmt.Println(s[start:end])
		start += size
		end += size
	}
	fmt.Println()
}
