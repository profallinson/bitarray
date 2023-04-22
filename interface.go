package bitarray

type BitArray interface {
	Write(uint64, bool)
	Read(uint64) bool
	Size() uint64

	GetBlock(int) bitBlock
	SetBlock(int, bitBlock)

	Equal(BitArray) bool
	Difference(BitArray) BitArray
	DistanceInt(BitArray) uint64
	Distance(BitArray) float64
	Norm() uint64
	Complement() BitArray
	Copy() BitArray
	Union(BitArray) BitArray
	RemainderInt(BitArray) int
	Remainder(BitArray) float64
	OverlapInt(BitArray) int
	Overlap(BitArray) float64
	Between(BitArray, BitArray) bool
	Sparsity() float64

	Rotate(float64) BitArray
	Move(int, int) BitArray
	Clip(uint8, uint8, uint64) BitArray
	Contrast(uint8, uint8, uint64) BitArray
	Avg(uint64) BitArray
	Binarify(uint8, uint8) BitArray

	ToString() string
	ToBytes() []uint8
	ToUint64() []uint64
	ToBools() []bool
}
