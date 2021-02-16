package bitarray

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestBitBlock(t *testing.T) {

	BeforeEach(func() {

	})

	AfterEach(func() {

	})

	Describe("bitBlock", func() {

		It("should set the first bit to 1", func() {
			var b bitBlock
			b = b.write(0, true)
			AssertEqual(b.read(0), true)
		})

		It("should set the first bit to 1 then back to 0", func() {
			var b bitBlock
			b = b.write(0, true)
			b = b.write(0, false)
			AssertEqual(b.read(0), false)
		})

		It("should panic because of writing to an out of range index", func() {
			defer func() {
				if r := recover(); r != nil {
					AssertEqual(true, true)
				}
			}()
			var b bitBlock
			b = b.write(64, true)
			AssertEqual(false, true)
		})

		It("should panic because of reading from an out of range index", func() {
			defer func() {
				if r := recover(); r != nil {
					AssertEqual(true, true)
				}
			}()
			var b bitBlock
			b.read(64)
			AssertEqual(false, true)
		})

		It("should set the third bit to 1 and return 8", func() {
			var b bitBlock
			b = b.write(3, true)
			AssertEqual(uint64(b), uint64(8))
		})

		It("should set the fourth bit to 1 and return 16", func() {
			var b bitBlock
			b = b.write(4, true)
			AssertEqual(b.ToNumber(), uint64(16))
		})

		It("should set the last bit to 1 and return a string value", func() {
			var b bitBlock
			b = b.write(0, true)
			AssertEqual(b.ToString(), "0000000000000000000000000000000000000000000000000000000000000001")
		})

		It("should return a string representation of 1,000,000", func() {
			var b bitBlock = 1000000
			AssertEqual(b.ToString(), "0000000000000000000000000000000000000000000011110100001001000000")
		})

	})

	Report(t)
}
