package bitarray

import (
	"fmt"
	"strings"
)

const (
	PRINT_TRUE  = "X"
	PRINT_FALSE = "."
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

// Returns a string of given width and height determined the size of the BitArray.
func Print(a BitArray, width int) string {
	if a == nil {
		return ""
	}
	ret := ""
	row := width
	for y := uint64(0); y < a.Size(); y++ {
		if a.Read(y) {
			ret += fmt.Sprint(PRINT_TRUE)
		} else {
			ret += fmt.Sprint(PRINT_FALSE)
		}
		row--
		if row == 0 {
			ret += fmt.Sprint("\n")
			row = width
		}
	}
	return ret
}

func appendEmptyRows(rows []string, size int, width int) []string {
	extra := strings.Repeat(" ", width)
	for i := len(rows); i < size; i++ {
		rows = append(rows, extra)
	}
	return rows
}

// Returns a sting of given width with the BitArrays side by side.
func PrintSideBySide(a BitArray, b BitArray, width int) string {
	ret := ""
	left := strings.Split(Print(a, width), "\n")
	right := strings.Split(Print(b, width), "\n")
	left = left[:len(left)-1]    // the last item is an empty string so remove it.
	right = right[:len(right)-1] // the last item is an empty string so remove it.
	rows := len(left)
	if rows > len(right) {
		right = appendEmptyRows(right, rows, width)
	} else if rows < len(right) {
		rows = len(right)
		left = appendEmptyRows(left, rows, width)
	}
	for i := 0; i < rows; i++ {
		ret += left[i] + " " + right[i] + "\n"
	}
	return ret
}
