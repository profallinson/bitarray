package bitarray

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestOperations(t *testing.T) {

	BeforeEach(func() {

	})

	AfterEach(func() {

	})

	Describe("Equals()", func() {

		It("should return 'true' as they are the same", func() {
			a := NewBitArrayFromString("1010101010101010101010")
			b := NewBitArrayFromString("1010101010101010101010")
			AssertEqual(a.Equal(b), true)
		})

		It("should return 'false' as they are NOT the same", func() {
			a := NewBitArrayFromString("1010101010101010101010")
			b := NewBitArrayFromString("0010101010101010101010")
			AssertEqual(a.Equal(b), false)
		})

		It("should return 'false' as they are different sizes (in multiples of 64)", func() {
			a := NewBitArrayOfLength(64)
			b := NewBitArrayOfLength(128)
			AssertEqual(a.Equal(b), false)
		})

	})

	Describe("Difference()", func() {

		It("should return all '1's", func() {
			a := NewBitArrayFromString("10001")
			b := NewBitArrayFromString("01110")
			AssertEqual(a.Difference(b).ToString()[:5], "11111")
		})

		It("should return '1's at the start and end", func() {
			a := NewBitArrayFromString("11111")
			b := NewBitArrayFromString("01110")
			AssertEqual(a.Difference(b).ToString()[:5], "10001")
		})

		It("should 'panic()' as they are different sizes (in multiples of 64)", func() {
			defer func() {
				if r := recover(); r != nil {
					AssertEqual(true, true)
				}
			}()
			a := NewBitArrayOfLength(64)
			b := NewBitArrayOfLength(65)
			a.Difference(b)
			AssertEqual(true, false)
		})

	})

	Describe("Norm()", func() {

		It("should return 1", func() {
			a := NewBitArrayFromString("10000")
			AssertEqual(a.Norm(), uint64(1))
		})

		It("should return 2", func() {
			a := NewBitArrayFromString("10001")
			AssertEqual(a.Norm(), uint64(2))
		})

		It("should return 3", func() {
			a := NewBitArrayFromString("10101")
			AssertEqual(a.Norm(), uint64(3))
		})

		It("should return 5", func() {
			a := NewBitArrayFromString("11111")
			AssertEqual(a.Norm(), uint64(5))
		})

	})

	Describe("Distance()", func() {

		It("should return 0", func() {
			a := NewBitArrayFromString("00001")
			b := NewBitArrayFromString("00001")
			AssertEqual(a.Distance(b), uint64(0))
		})

		It("should return 1", func() {
			a := NewBitArrayFromString("00001")
			b := NewBitArrayFromString("00000")
			AssertEqual(a.Distance(b), uint64(1))
		})

		It("should return 2", func() {
			a := NewBitArrayFromString("00001")
			b := NewBitArrayFromString("10000")
			AssertEqual(a.Distance(b), uint64(2))
		})

		It("should return 5", func() {
			a := NewBitArrayFromString("10001")
			b := NewBitArrayFromString("01110")
			AssertEqual(a.Distance(b), uint64(5))
		})

	})

	Describe("Copy()", func() {

		It("should return a new array which is not affected by changes to the original", func() {
			a := NewBitArrayFromString("10000")
			b := a.Copy()
			a.Write(0, false)
			AssertEqual(a.Read(0), false)
			AssertEqual(b.Read(0), true)
		})

	})

	Describe("Complement()", func() {

		It("should return a new array which is the opposite of the original", func() {
			a := NewBitArrayFromString("10000")
			b := a.Complement()
			AssertEqual(b.ToString()[:5], "01111")
		})

	})

	Describe("Union()", func() {

		It("should 'panic()' as they are different sizes (in multiples of 64)", func() {
			defer func() {
				if r := recover(); r != nil {
					AssertEqual(true, true)
				}
			}()
			a := NewBitArrayOfLength(64)
			b := NewBitArrayOfLength(65)
			a.Union(b)
			AssertEqual(true, false)
		})

		It("should return a new array of '10001'", func() {
			a := NewBitArrayFromString("10000")
			b := NewBitArrayFromString("00001")
			c := a.Union(b)
			AssertEqual(c.ToString()[:5], "10001")
		})

		It("should return a new array of '10101'", func() {
			a := NewBitArrayFromString("10100")
			b := NewBitArrayFromString("00101")
			c := a.Union(b)
			AssertEqual(c.ToString()[:5], "10101")
		})

	})

	Describe("Remainder()", func() {

		It("should 'panic()' as they are different sizes (in multiples of 64)", func() {
			defer func() {
				if r := recover(); r != nil {
					AssertEqual(true, true)
				}
			}()
			a := NewBitArrayOfLength(64)
			b := NewBitArrayOfLength(65)
			a.RemainderInt(b)
			AssertEqual(true, false)
		})

		It("should ", func() {
			a := NewBitArrayFromString("10000")
			b := NewBitArrayFromString("01001")
			AssertEqual(a.RemainderInt(b), 1)
		})

		It("should ", func() {
			a := NewBitArrayFromString("11111")
			b := NewBitArrayFromString("10001")
			AssertEqual(a.RemainderInt(b), 3)
		})

		It("should 'panic()' as they are different sizes (in multiples of 64)", func() {
			defer func() {
				if r := recover(); r != nil {
					AssertEqual(true, true)
				}
			}()
			a := NewBitArrayOfLength(64)
			b := NewBitArrayOfLength(65)
			a.Remainder(b)
			AssertEqual(true, false)
		})

		It("should ", func() {
			a := NewBitArrayFromString("10000")
			b := NewBitArrayFromString("01001")
			AssertEqual(a.Remainder(b), float32(1))
		})

		It("should ", func() {
			a := NewBitArrayFromString("11111")
			b := NewBitArrayFromString("10001")
			AssertEqual(a.Remainder(b), float32(0.6))
		})

	})

	Describe("Overlap()", func() {

		It("should 'panic()' as they are different sizes (in multiples of 64)", func() {
			defer func() {
				if r := recover(); r != nil {
					AssertEqual(true, true)
				}
			}()
			a := NewBitArrayOfLength(64)
			b := NewBitArrayOfLength(65)
			a.OverlapInt(b)
			AssertEqual(true, false)
		})

		It("should return 2", func() {
			a := NewBitArrayFromString("10010")
			b := NewBitArrayFromString("10110")
			AssertEqual(a.OverlapInt(b), 2)
		})

		It("should return 0", func() {
			a := NewBitArrayFromString("11110")
			b := NewBitArrayFromString("00001")
			AssertEqual(a.OverlapInt(b), 0)
		})

		It("should 'panic()' as they are different sizes (in multiples of 64)", func() {
			defer func() {
				if r := recover(); r != nil {
					AssertEqual(true, true)
				}
			}()
			a := NewBitArrayOfLength(64)
			b := NewBitArrayOfLength(65)
			a.Overlap(b)
			AssertEqual(true, false)
		})

		It("should return 0.03125", func() {
			a := NewBitArrayFromString("10010")
			b := NewBitArrayFromString("10110")
			AssertEqual(a.Overlap(b), float32(0.03125))
		})

		It("should return 0.0", func() {
			a := NewBitArrayFromString("11110")
			b := NewBitArrayFromString("00001")
			AssertEqual(a.Overlap(b), float32(0.0))
		})

	})

	Describe("Betweenness()", func() {

		It("should return 'false' as 'a' is a different length", func() {
			a := NewBitArrayOfLength(128)
			b := NewBitArrayOfLength(64)
			c := NewBitArrayOfLength(64)
			AssertEqual(b.Between(a, c), false)
		})

		It("should return 'false' as 'b' is a different length", func() {
			a := NewBitArrayOfLength(64)
			b := NewBitArrayOfLength(128)
			c := NewBitArrayOfLength(64)
			AssertEqual(b.Between(a, c), false)
		})

		It("should return 'false' as 'c' is a different length", func() {
			a := NewBitArrayOfLength(64)
			b := NewBitArrayOfLength(128)
			c := NewBitArrayOfLength(64)
			AssertEqual(b.Between(a, c), false)
		})

		It("should return 'true' because?", func() {
			a := NewBitArrayFromString("00111")
			b := NewBitArrayFromString("01110")
			c := NewBitArrayFromString("01010")
			AssertEqual(b.Between(a, c), true)
		})

	})

	Describe("Sparsity()", func() {

		It("should return 0", func() {
			a := NewBitArrayFromString("0")
			AssertEqual(a.Sparsity(), 0.0)
		})

		It("should return 1.5625", func() {
			a := NewBitArrayFromString("1")
			AssertEqual(a.Sparsity(), 1.5625)
		})

		It("should return 100.0", func() {
			a := NewBitArray(1).Complement()
			AssertEqual(a.Sparsity(), 100.0)
		})

	})

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

	Report(t)
}
