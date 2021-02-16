package bitarray

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestBitArray(t *testing.T) {

	BeforeEach(func() {

	})

	AfterEach(func() {

	})

	Describe("BitArray", func() {

		It("should set the 0 bit to 'true'", func() {
			this := NewBitArray(1)
			this.Write(0, true)
			AssertEqual(this.Read(0), true)
		})

		It("should print a string with the first bit as 1", func() {
			this := NewBitArray(1)
			this.Write(0, true)
			AssertEqual(this.ToString(), "1000000000000000000000000000000000000000000000000000000000000000")
		})

		It("should print a string with the last bit as 1", func() {
			this := NewBitArray(1)
			this.Write(63, true)
			AssertEqual(this.ToString(), "0000000000000000000000000000000000000000000000000000000000000001")
		})

		It("should set the last bit to 'true' when 2 blocks are used", func() {
			this := NewBitArray(2)
			this.Write(127, true)
			AssertEqual(this.Read(127), true)
		})

		It("should set the last bit to 'true' when 3 blocks are used", func() {
			this := NewBitArray(3)
			this.Write(191, true)
			AssertEqual(this.Read(191), true)
		})

		It("should have the first and last bits set to 1, all else 0", func() {
			this := NewBitArray(3)
			this.Write(0, true)
			this.Write(191, true)
			AssertEqual(this.ToString(), "100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001")
		})

		It("should have the first and last bits of each block set to 1, all else 0", func() {
			this := NewBitArray(3)
			this.Write(0, true)
			this.Write(63, true)
			this.Write(64, true)
			this.Write(127, true)
			this.Write(128, true)
			this.Write(191, true)
			AssertEqual(this.ToString(), "100000000000000000000000000000000000000000000000000000000000000110000000000000000000000000000000000000000000000000000000000000011000000000000000000000000000000000000000000000000000000000000001")
		})

		It("should return the number 1", func() {
			this := NewBitArray(1)
			this.Write(0, true)
			AssertEqual(this.ToUint64()[0], uint64(1))
		})

		It("should return the numbers 1, 0, 9223372036854775808", func() {
			this := NewBitArray(3)
			this.Write(0, true)
			this.Write(191, true)
			AssertEqual(this.ToUint64()[0], uint64(1))
			AssertEqual(this.ToUint64()[1], uint64(0))
			AssertEqual(this.ToUint64()[2], uint64(9223372036854775808))
		})

		It("should return the all numbers as 1", func() {
			this := NewBitArray(1)
			this.Write(0, true)
			this.Write(8, true)
			this.Write(16, true)
			this.Write(24, true)
			this.Write(32, true)
			this.Write(40, true)
			this.Write(48, true)
			this.Write(56, true)
			AssertEqual(this.ToBytes()[0], uint8(1))
			AssertEqual(this.ToBytes()[1], uint8(1))
			AssertEqual(this.ToBytes()[2], uint8(1))
			AssertEqual(this.ToBytes()[3], uint8(1))
			AssertEqual(this.ToBytes()[4], uint8(1))
			AssertEqual(this.ToBytes()[5], uint8(1))
			AssertEqual(this.ToBytes()[6], uint8(1))
			AssertEqual(this.ToBytes()[7], uint8(1))
		})

	})

	Report(t)
}
