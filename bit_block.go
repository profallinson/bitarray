package bitarray

import (
	"fmt"
)

const bitBlockSize = 64

type bitBlock uint64

func (this bitBlock) write(position uint64, value bool) bitBlock {
	if position >= bitBlockSize {
		panic("Index out of range.")
	}
	if value {
		return this | bitBlock(1<<position)
	}
	return this & ^bitBlock(1<<position)
}

func (this bitBlock) read(position uint64) bool {
	if position >= bitBlockSize {
		panic("Index out of range.")
	}
	return this&bitBlock(1<<position) != 0
}

func (this bitBlock) ToString() string {
	return fmt.Sprintf(fmt.Sprintf("%%0%db", bitBlockSize), uint64(this))
}

func (this bitBlock) ToNumber() uint64 {
	return uint64(this)
}

// Returns an array of booleans representing the block.
func (this bitBlock) ToBools() []bool {
	r := make([]bool, bitBlockSize)
	var i uint64
	for i = 0; i < bitBlockSize; i++ {
		r[i] = this.read(i)
	}
	return r
}
