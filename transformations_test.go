package bitarray

import (
	"fmt"
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestTransformations(t *testing.T) {

	Describe("Rotate()", func() {

		// 10011000
		// 00111100
		// 01000010
		// 10000001
		// 10000001
		// 01000010
		// 00100100
		// 00011000
		source := "1001100000111100010000101000000110000001010000100010010000011000"

		It("should rotate the bit array by 90 deg", func() {
			a := NewBitArrayFromString(source)
			n := a.Rotate(90)
			PrintMatrix(90, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0001100100100100010000101000001110000011010000100010010000011000")
		})

		It("should rotate the bit array by -90 deg", func() {
			a := NewBitArrayFromString(source)
			n := a.Rotate(-90)
			PrintMatrix(-90, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0001100000100100010000101100000111000001010000100010010010011000")
		})

		It("should rotate the bit array by 180 deg", func() {
			a := NewBitArrayFromString(source)
			n := a.Rotate(180)
			PrintMatrix(180, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0001100000100100010000101000000110000001010000100011110000011001")
		})

		It("should rotate the bit array by 270 deg", func() {
			a := NewBitArrayFromString(source)
			n := a.Rotate(270)
			PrintMatrix(270, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0001100000100100010000101100000111000001010000100010010010011000")
		})

		It("should rotate the bit array by 360 deg", func() {
			a := NewBitArrayFromString(source)
			n := a.Rotate(360)
			PrintMatrix(360, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), a.ToString())
		})

		It("should rotate the bit array by 720 deg", func() {
			a := NewBitArrayFromString(source)
			n := a.Rotate(720)
			PrintMatrix(720, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), a.ToString())
		})

		It("should rotate the bit array by 45 deg", func() {
			a := NewBitArrayFromString(source)
			n := a.Rotate(45)
			PrintMatrix(45, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0000000001010110000000100000001000000010000000000101101000000000")
		})

	})

	Describe("Move()", func() {

		// 10011000
		// 00111100
		// 01000010
		// 10000001
		// 10000001
		// 01000010
		// 00100100
		// 00011000
		source := "1001100000111100010000101000000110000001010000100010010000011000"

		It("should move the bit array by 1 row up", func() {
			a := NewBitArrayFromString(source)
			n := a.Move(0, 1)
			PrintMatrix(1, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0011110001000010100000011000000101000010001001000001100000000000")
		})

		It("should move the bit array by 1 row down", func() {
			a := NewBitArrayFromString(source)
			n := a.Move(0, -1)
			PrintMatrix(-1, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0000000010011000001111000100001010000001100000010100001000100100")
		})

		It("should move the bit array by 4 rows up", func() {
			a := NewBitArrayFromString(source)
			n := a.Move(0, 4)
			PrintMatrix(4, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "1000000101000010001001000001100000000000000000000000000000000000")
		})

		It("should move the bit array by 4 rows down", func() {
			a := NewBitArrayFromString(source)
			n := a.Move(0, -4)
			PrintMatrix(-4, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0000000000000000000000000000000010011000001111000100001010000001")
		})

		It("should move the bit array by 2 rows right", func() {
			a := NewBitArrayFromString(source)
			n := a.Move(2, 0)
			PrintMatrix(2, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0010011000001111000100000010000000100000000100000000100100000110")
		})

		It("should move the bit array by 2 rows left", func() {
			a := NewBitArrayFromString(source)
			n := a.Move(-2, 0)
			PrintMatrix(-2, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0110000011110000000010000000010000000100000010001001000001100000")
		})

		It("should move the bit array by 5 rows right", func() {
			a := NewBitArrayFromString(source)
			n := a.Move(5, 0)
			PrintMatrix(5, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0000010000000001000000100000010000000100000000100000000100000000")
		})

		It("should move the bit array by 5 rows left", func() {
			a := NewBitArrayFromString(source)
			n := a.Move(-5, 0)
			PrintMatrix(-5, n.ToString(), uint64(8))
			AssertEqual(n.ToString(), "0000000010000000010000000010000000100000010000001000000000000000")
		})
	})

	Describe("Clip()", func() {

		It("should return a byte array of 0s with the last five set to 255s", func() {
			source := []byte{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150}
			a := NewBitArrayFromBytes(source).Clip(100, 255, 1).ToBytes()
			test := true
			for i, _ := range a {
				switch {
				case i >= 10 && i <= 15 && a[i] != byte(a[i]):
					test = false
				case i <= 9 && a[i] != byte(0):
					test = false
				}
				// fmt.Println(source[i], a[i])
			}
			AssertEqual(test, true)
		})

		It("should return a byte array of 0s with just two set to 255s", func() {
			source := []byte{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150}
			a := NewBitArrayFromBytes(source).Clip(100, 110, 1).ToBytes()
			test := true
			for i, _ := range a {
				switch {
				case i < 10 && i > 11 && a[i] != byte(0):
					test = false
				case (i == 10 || i == 11) && a[i] != byte(a[i]):
					test = false
				}
				// fmt.Println(source[i], a[i])
			}
			AssertEqual(test, true)
		})

		It("should return a byte array of 0s with two sets of three set to 255s", func() {
			source := []byte{100, 100, 20, 30, 40, 100, 100, 100, 100, 90, 100, 110, 120, 130, 140, 150}
			a := NewBitArrayFromBytes(source).Clip(100, 255, 3).ToBytes()
			test := true
			for i, _ := range a {
				switch {
				case i >= 6 && i <= 8 && a[i] != byte(a[i]):
					test = false
				case i >= 12 && i <= 14 && a[i] != byte(a[i]):
					test = false
				case (i <= 5 || i >= 15 || (i >= 9 && i <= 11)) && a[i] != byte(0):
					test = false
				}
				// fmt.Println(source[i], a[i])
			}
			AssertEqual(test, true)
		})
	})

	Describe("Contrast()", func() {

		It("should return a byte array of 0s with the last five set to 255s", func() {
			source := []byte{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150}
			a := NewBitArrayFromBytes(source).Contrast(100, 255, 1).ToBytes()
			test := true
			for i, _ := range a {
				switch {
				case i >= 10 && i <= 15 && a[i] != byte(255):
					test = false
				case i <= 9 && a[i] != byte(0):
					test = false
				}
				// fmt.Println(source[i], a[i])
			}
			AssertEqual(test, true)
		})

		It("should return a byte array of 0s with just two set to 255s", func() {
			source := []byte{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150}
			a := NewBitArrayFromBytes(source).Contrast(100, 110, 1).ToBytes()
			test := true
			for i, _ := range a {
				switch {
				case i < 10 && i > 11 && a[i] != byte(0):
					test = false
				case (i == 10 || i == 11) && a[i] != byte(255):
					test = false
				}
				// fmt.Println(source[i], a[i])
			}
			AssertEqual(test, true)
		})

		It("should return a byte array of 0s with two sets of three set to 255s", func() {
			source := []byte{100, 100, 20, 30, 40, 100, 100, 100, 100, 90, 100, 110, 120, 130, 140, 150}
			a := NewBitArrayFromBytes(source).Contrast(100, 255, 3).ToBytes()
			test := true
			for i, _ := range a {
				switch {
				case i >= 6 && i <= 8 && a[i] != byte(255):
					test = false
				case i >= 12 && i <= 14 && a[i] != byte(255):
					test = false
				case (i <= 5 || i >= 15 || (i >= 9 && i <= 11)) && a[i] != byte(0):
					test = false
				}
				// fmt.Println(source[i], a[i])
			}
			AssertEqual(test, true)
		})

	})

	Describe("Scale()", func() {

		It("should return a length of 16", func() {
			source := make([]byte, 16)
			a := NewBitArrayFromBytes(source).Avg(1).ToBytes()
			AssertEqual(len(a), 16)
		})

		It("should return a length of 8", func() {
			source := make([]byte, 16)
			a := NewBitArrayFromBytes(source).Avg(2).ToBytes()
			AssertEqual(len(a), 8)
		})

		It("should return a length of 32", func() {
			source := make([]byte, 64)
			a := NewBitArrayFromBytes(source).Avg(4).ToBytes()
			AssertEqual(len(a), 16)
		})

		It("should return a length of 32", func() {
			source := []byte{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150}
			a := NewBitArrayFromBytes(source).Avg(2).ToBytes()
			test := true
			for range a {
				switch {
				case a[0] != byte(5):
					test = false
				case a[1] != byte(25):
					test = false
				case a[2] != byte(45):
					test = false
				case a[3] != byte(65):
					test = false
				case a[4] != byte(85):
					test = false
				case a[5] != byte(105):
					test = false
				case a[6] != byte(125):
					test = false
				case a[7] != byte(145):
					test = false
				}
			}
			AssertEqual(test, true)
		})

	})

	Describe("Binarify()", func() {

		It("should return ", func() {
			source := make([]byte, 64)
			for i, _ := range source {
				source[i] = byte(i + 50)
			}
			a := NewBitArrayFromBytes(source).Binarify(100, 255)
			AssertEqual(a.ToString(), "0000000000000000000000000000000000000000000000000011111111111111")
		})

	})

	Describe("Radius()", func() {

		It("should return the first four of the top four rows", func() {
			a := NewBitArrayFromString("1111111111111111000000000000000000000000000000001111111111111111")
			b := a.Radius(2, 2, 2)
			fmt.Println(Print(a, 8))
			fmt.Println(Print(b, 4))
			AssertEqual(b.ToString(), "1111111100000000000000000000000000000000000000000000000000000000")
		})

		It("should return the last four of bottom four rows", func() {
			a := NewBitArrayFromString("1111111111111111000000000000000000000000000000001111111111111111")
			b := a.Radius(6, 6, 2)
			fmt.Println(Print(a, 8))
			fmt.Println(Print(b, 4))
			AssertEqual(b.ToString(), "0000000011111111000000000000000000000000000000000000000000000000")
		})

	})

	Report(t)
}
