package bitarray

import (
	"fmt"
)

// Size is the number of 64 bit blocks to use.
func NewBitArray(blocks uint64) BitArray {
	return createBitArray(blocks)
}

// Size is the number of bits to use.
func NewBitArrayOfLength(length uint64) BitArray {
	size := uint64(length / bitBlockSize)
	// If the given string is not divisible by 'bitBlockSize' then increase the 'size' by 1.
	if length%bitBlockSize > 0 {
		size++
	}
	return createBitArray(size)
}

// TODO: Add distribution throughout the array.
// Returns an array of size 'blocks' initialized with the given 'sparsity'.
func NewBitArrayOfSparsity(length uint64, sparsity float32) BitArray {
	this := NewBitArrayOfLength(length)
	space := uint64(float32(length) / 100 * sparsity)
	fmt.Println(this.Size(), space)
	for i := uint64(0); i < space; i++ {
		this.Write(i, true)
	}
	return this
}

// Given a string of '0's and '1's returns a populated bit array.
func NewBitArrayFromString(str string) BitArray {
	size := uint64(len(str) / bitBlockSize)
	// If the given string is not divisible by 'bitBlockSize' then increase the 'size' by 1.
	if len(str)%bitBlockSize > 0 {
		size++
	}
	this := createBitArray(size)
	for i := range str {
		this.Write(uint64(i), getBool(str[i]))
	}
	return this
}

// Given an array of 'uint8' returns a populated bit array.
func NewBitArrayFromBytes(nums []uint8) BitArray {
	// Make sure the length of 'nums' is a multiple of 8.
	for i := 0; i < len(nums)%8; i++ {
		nums = append(nums, 0)
	}
	// Create an array on uint64s to pack the uint8s into.
	pack := make([]uint64, len(nums)/8)
	p := 0
	for i := 0; i < len(nums); i += 8 {
		pack[p] = uint64(nums[i]) | uint64(nums[i+1])<<8 | uint64(nums[i+2])<<16 | uint64(nums[i+3])<<24 | uint64(nums[i+4])<<32 | uint64(nums[i+5])<<40 | uint64(nums[i+6])<<48 | uint64(nums[i+7])<<56
		p++
	}
	return NewBitArrayFromUint64(pack)
}

// Given an array of 'uint64' returns a populated bit array.
func NewBitArrayFromUint64(nums []uint64) BitArray {
	size := uint64(len(nums))
	this := createBitArray(size)
	block := 0
	for n := 0; n < len(nums); n++ {
		for i := 0; i < 64; i++ {
			this.Write(uint64(block*bitBlockSize+i), uint64(nums[n])&uint64(1<<uint64(i)) != 0)
		}
		block++
	}
	return this
}

// Given an array of booleans returns a populated bit array.
func NewBitArrayFromBools(bools []bool) BitArray {
	size := uint64(len(bools) / bitBlockSize)
	// If the given string is not divisible by 'bitBlockSize' then increase the 'size' by 1.
	if len(bools)%bitBlockSize > 0 {
		size++
	}
	this := createBitArray(size)
	for i := range bools {
		this.Write(uint64(i), bools[i])
	}
	return this
}

// Returns 'true' for '1' or 'false' for all others.
func getBool(b byte) bool {
	switch b {
	case '1':
		return true
	}
	return false
}
