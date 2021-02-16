package bitarray

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestNewBitArray(t *testing.T) {

	BeforeEach(func() {

	})

	AfterEach(func() {

	})

	Describe("NewBitArray()", func() {

		It("should set the 0 bit to 'true'", func() {
			this := NewBitArray(1)
			this.Write(0, true)
			AssertEqual(this.Read(0), true)
		})

	})

	Describe("NewBitArrayFromString()", func() {

		It("should return index 50 as a 'true'", func() {
			this := NewBitArray(1)
			this.Write(50, true)
			str := this.ToString()
			test := NewBitArrayFromString(str)
			AssertEqual(test.Read(50), true)
		})

		It("should return index's 63, 64 and 191 as true", func() {
			this := NewBitArray(3)
			this.Write(63, true)
			this.Write(64, true)
			this.Write(191, true)
			str := this.ToString()
			test := NewBitArrayFromString(str)
			AssertEqual(test.Read(63), true)
			AssertEqual(test.Read(64), true)
			AssertEqual(test.Read(191), true)
		})

		It("should create a BitArray for the given size input", func() {
			this := NewBitArrayFromString("1010101010101010101010")
			AssertEqual(this.ToUint64()[0], uint64(1398101))
		})

	})

	Describe("NewBitArrayFromBytes()", func() {

		It("should return uint64 1 from 8 uint8 values", func() {
			nums := []uint8{1, 0, 0, 0, 0, 0, 0, 0}
			test := NewBitArrayFromBytes(nums)
			AssertEqual(test.ToBytes()[0], uint8(1))
		})

		It("should return uint64 1 from 8 uint8 values after filling in missing 0", func() {
			nums := []uint8{1}
			test := NewBitArrayFromBytes(nums)
			AssertEqual(test.ToBytes()[0], uint8(1))
		})

		It("should return uint8 values in the same order given up to 64 bits", func() {
			nums := []uint8{128, 255, 0, 0, 0, 0, 0, 69}
			test := NewBitArrayFromBytes(nums)
			AssertEqual(test.ToBytes()[0], uint8(128))
			AssertEqual(test.ToBytes()[1], uint8(255))
			AssertEqual(test.ToBytes()[7], uint8(69))
		})

		It("should return uint8 values in the same order given up to 128 bits", func() {
			nums := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8}
			test := NewBitArrayFromBytes(nums)
			AssertEqual(test.ToBytes()[8], uint8(1))
			AssertEqual(test.ToBytes()[9], uint8(2))
			AssertEqual(test.ToBytes()[15], uint8(8))
		})

		It("should return the same number after creation", func() {
			this := NewBitArray(1)
			this.Write(0, true)
			this.Write(8, true)
			this.Write(16, true)
			this.Write(24, true)
			this.Write(32, true)
			this.Write(40, true)
			this.Write(48, true)
			this.Write(56, true)
			nums := this.ToBytes()

			test := NewBitArrayFromBytes(nums)

			AssertEqual(test.ToBytes()[0], nums[0])
			AssertEqual(test.ToBytes()[1], nums[1])
			AssertEqual(test.ToBytes()[2], nums[2])
			AssertEqual(test.ToBytes()[3], nums[3])
			AssertEqual(test.ToBytes()[4], nums[4])
			AssertEqual(test.ToBytes()[5], nums[5])
			AssertEqual(test.ToBytes()[6], nums[6])
			AssertEqual(test.ToBytes()[7], nums[7])
		})

	})

	Describe("NewBitArrayFromUint64()", func() {

		It("should return the same number after creation", func() {
			this := NewBitArray(1)
			this.Write(0, true)
			this.Write(63, true)
			nums := this.ToUint64()

			test := NewBitArrayFromUint64(nums)

			AssertEqual(test.ToUint64()[0], nums[0])
		})

		It("should return the same number after nested creation", func() {
			this := NewBitArray(1)
			this.Write(0, true)
			this.Write(63, true)
			nums := this.ToUint64()

			test := NewBitArrayFromUint64(NewBitArrayFromUint64(nums).ToUint64())

			AssertEqual(test.ToUint64()[0], nums[0])
		})

		It("should return the same 3 numbers in the same order after creation", func() {
			this := NewBitArray(3)
			this.Write(63, true)
			this.Write(64, true)
			this.Write(191, true)
			nums := this.ToUint64()

			test := NewBitArrayFromUint64(nums).ToUint64()

			AssertEqual(test[0], nums[0])
			AssertEqual(test[1], nums[1])
			AssertEqual(test[2], nums[2])
		})

	})

	Describe("NewBitArrayFromBools()", func() {

		It("should return the same 3 first booleans in the same order after creation", func() {
			this := NewBitArrayFromBools([]bool{true, false, true, false, true, false, true, false, true, false})
			bools := this.ToBools()

			test := NewBitArrayFromBools(bools).ToBools()

			AssertEqual(test[0], bools[0])
			AssertEqual(test[1], bools[1])
			AssertEqual(test[2], bools[2])
		})

	})

	Describe("NewBitArrayOfLength()", func() {

		It("should return a new BitArray 64 bits long", func() {
			this := NewBitArrayOfLength(uint64(64))
			AssertEqual(this.Size(), uint64(64))
		})

		It("should return a new BitArray 128 bits long", func() {
			this := NewBitArrayOfLength(uint64(65))
			AssertEqual(this.Size(), uint64(128))
		})

	})

	Describe("NewBitArrayOfSparsity()", func() {

		It("should return a Norm() of 12", func() {
			this := NewBitArrayOfSparsity(64, 20)
			AssertEqual(this.Norm(), uint64(12))
		})

		It("should return a Norm() of 25", func() {
			this := NewBitArrayOfSparsity(64, 40)
			AssertEqual(this.Norm(), uint64(25))
		})

		It("should return a Norm() of 38", func() {
			this := NewBitArrayOfSparsity(64, 60)
			AssertEqual(this.Norm(), uint64(38))
		})

		It("should return a Norm() of 51", func() {
			this := NewBitArrayOfSparsity(64, 80)
			AssertEqual(this.Norm(), uint64(51))
		})

		It("should return a Norm() of 64", func() {
			this := NewBitArrayOfSparsity(64, 100)
			AssertEqual(this.Norm(), uint64(64))
		})

	})

	Report(t)
}
