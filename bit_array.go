package bitarray

type bitArray struct {
	blocks []bitBlock
}

// Write a bit at the given location.
func (this *bitArray) Write(index uint64, value bool) {
	id, pos := this.getBitBlockAddress(index)
	this.blocks[id] = this.blocks[id].write(pos, value)
}

// Read the bit from the given location.
func (this *bitArray) Read(index uint64) bool {
	id, pos := this.getBitBlockAddress(index)
	return this.blocks[id].read(pos)
}

// Returns a string of '0's '1's representing the bit array.
func (this *bitArray) ToString() string {
	s := ""
	for i, _ := range this.blocks {
		s += this.reverse(this.blocks[i].ToString())
	}
	return s
}

func (this *bitArray) GetBlock(i int) bitBlock {
	return this.blocks[i]
}

func (this *bitArray) SetBlock(i int, block bitBlock) {
	this.blocks[i] = block
}

// Returns an array of uint8 representing the bit array.
func (this *bitArray) ToBytes() []uint8 {
	r := make([]uint8, len(this.blocks)*8)
	for i := 0; i < len(this.blocks); i++ {
		num := this.blocks[i].ToNumber()
		r[i*8+0] = uint8(num >> 0)
		r[i*8+1] = uint8(num >> 8)
		r[i*8+2] = uint8(num >> 16)
		r[i*8+3] = uint8(num >> 24)
		r[i*8+4] = uint8(num >> 32)
		r[i*8+5] = uint8(num >> 40)
		r[i*8+6] = uint8(num >> 48)
		r[i*8+7] = uint8(num >> 56)
	}
	return r
}

// Returns an array of uint64 representing the bit array.
func (this *bitArray) ToUint64() []uint64 {
	r := make([]uint64, len(this.blocks))
	for i, _ := range this.blocks {
		r[i] = this.blocks[i].ToNumber()
	}
	return r
}

// Returns an array of booleans representing the bit array.
func (this *bitArray) ToBools() []bool {
	r := make([]bool, 0)
	for i, _ := range this.blocks {
		r = append(r, this.blocks[i].ToBools()...)
	}
	return r
}

// Returns the size of the array in bits.
func (this *bitArray) Size() uint64 {
	return uint64(len(this.blocks) * bitBlockSize)
}

// Returns the given string reversed so the first is last and the last is first, etc.
func (this *bitArray) reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func (this *bitArray) getBitBlockAddress(index uint64) (uint64, uint64) {
	id := (index / bitBlockSize)
	pos := index - (id * bitBlockSize)
	return id, pos
}

// Size is the number of 64 bit blocks to use.
func createBitArray(blocks uint64) *bitArray {
	this := &bitArray{}
	this.blocks = make([]bitBlock, blocks)
	return this
}
