package bitarray

type BitArray interface {
	Write(uint64, bool)
	Read(uint64) bool
	Size() uint64

	GetBlock(int) bitBlock
	SetBlock(int, bitBlock)

	Equal(BitArray) bool
	Difference(BitArray) BitArray
	Distance(BitArray) uint64
	Norm() uint64
	Complement() BitArray
	Copy() BitArray
	Union(BitArray) BitArray
	RemainderInt(BitArray) int
	Remainder(BitArray) float32
	OverlapInt(BitArray) int
	Overlap(BitArray) float32
	Between(BitArray, BitArray) bool
	Sparsity() float64
	Rotate(float64) BitArray
	Move(int, int) BitArray
	Contrast(uint8, uint8, uint64) BitArray

	ToString() string
	ToBytes() []uint8
	ToUint64() []uint64
	ToBools() []bool
}
