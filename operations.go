package bitarray

import (
	"fmt"
	"math"
	"math/bits"
)

// Returns 'true' if the given array is a bit match.
func (this *bitArray) Equal(a BitArray) bool {
	if this.Size() != a.Size() {
		return false
	}
	for i := range this.blocks {
		if this.GetBlock(i) != a.GetBlock(i) {
			return false
		}
	}
	return true
}

// Returns a new array with 'true' where 'a' and 'b' differ and 'false' elsewhere.
func (this *bitArray) Difference(a BitArray) BitArray {
	if this.Size() != a.Size() {
		panic("BitArrays MUST be the same length.")
	}
	c := NewBitArrayOfLength(this.Size())
	for i := range this.blocks {
		c.SetBlock(i, this.GetBlock(i)^a.GetBlock(i))
	}
	return c
}

// Returns the bit count of 'true' values.
func (this *bitArray) Norm() uint64 {
	n := uint64(0)
	for i := range this.blocks {
		n += uint64(bits.OnesCount64(uint64(this.GetBlock(i))))
	}
	return n
}

// Returns the bit count of the distance between 'a' and 'b'.
func (this *bitArray) Distance(a BitArray) uint64 {
	return this.Difference(a).Norm()
}

// Returns a deep copy of the given array.
func (this *bitArray) Copy() BitArray {
	c := NewBitArrayOfLength(this.Size())
	for i := range this.blocks {
		c.SetBlock(i, this.GetBlock(i))
	}
	return c
}

// Returns a new array as the opposite of the given array values.
func (this *bitArray) Complement() BitArray {
	c := NewBitArrayOfLength(this.Size())
	for i := range this.blocks {
		c.SetBlock(i, ^this.GetBlock(i))
	}
	return c
}

// Returns a new array where the 'true' values of both 'a' and 'b' are set 'true' with 'false' elsewhere.
func (this *bitArray) Union(a BitArray) BitArray {
	if this.Size() != a.Size() {
		panic("BitArrays MUST be the same length.")
	}
	c := NewBitArrayOfLength(this.Size())
	for i := range this.blocks {
		c.SetBlock(i, this.GetBlock(i)|a.GetBlock(i))
	}
	return c
}

// Returns the number of 'true' values in 'this' that are not contained in 'a'.
func (this *bitArray) remainderInt(a BitArray) (int, int) {
	if this.Size() != a.Size() {
		panic("BitArrays MUST be the same length.")
	}
	t := 0 // Count of 'true' values in 'a'.
	c := 0 // Count of 'true' values in 'a' that are not in the same position in 'b'.
	var i uint64
	for i = 0; i < this.Size(); i++ {
		if this.Read(i) {
			t++
		}
		if this.Read(i) && !a.Read(i) {
			c++
		}
	}
	return t, c
}

// Returns the number of 'true' values in 'a' that are not contained in 'b'.
func (this *bitArray) RemainderInt(a BitArray) int {
	_, c := this.remainderInt(a)
	return c
}

// Returns the percentage of 'true' values in 'a' that are not contained in 'b'.
func (this *bitArray) Remainder(a BitArray) float32 {
	t, c := this.remainderInt(a)
	return float32(c) / float32(t)
}

// Returns the number of the values that match between 'a' and 'b'.
func (this *bitArray) overlapInt(a BitArray) (int, int) {
	if this.Size() != a.Size() {
		panic("BitArrays MUST be the same length.")
	}
	o := 0
	for i := range this.blocks {
		o += bits.OnesCount64(uint64(this.GetBlock(i) & a.GetBlock(i)))
	}
	return int(this.Size()), o
}

// Returns the number of the values that match between 'a' and 'b'.
func (this *bitArray) OverlapInt(a BitArray) int {
	_, c := this.overlapInt(a)
	return c
}

// Returns the percentage of the values that match between 'a' and 'b'.
func (this *bitArray) Overlap(a BitArray) float32 {
	s, o := this.overlapInt(a)
	return float32(o) / float32(s)
}

// Returns 'true' if the current BitArray is found between 'a' and 'b'.
func (this *bitArray) Between(a BitArray, b BitArray) bool {
	if this.Size() != a.Size() || a.Size() != b.Size() {
		return false
	}
	return a.Distance(b) == a.Distance(this)+this.Distance(b)
}

// Returns the percentage of 'true' bits.
func (this *bitArray) Sparsity() float64 {
	return float64(this.Norm()) / float64(this.Size()) * 100
}

// If the bit array length can be squared to an `uint64` then it's content are rotated by the given `deg`.
func (this *bitArray) Rotate(deg float64) BitArray {
	sqrt := math.Sqrt(float64(this.Size()))
	size := uint64(sqrt)
	// halfSize := sqrt/2
	if sqrt != float64(size) {
		panic("Will only rotate arrays which can square to a `uint64`.")
	}
	fmt.Println(deg)
	n := NewBitArrayOfLength(this.Size())
	for w := uint64(0); w < size*size; w += size {
		oldY := float64(w/size)+1
		for oldI := w; oldI < w+size; oldI++ {
			oldX := float64(oldI - w)+1
			sin, cos := math.Sincos(deg*math.Pi/180)
			newX := math.Abs(math.Round((cos * oldX) + (sin * oldY)))
			newY := math.Abs(math.Round((cos * oldY) - (sin * oldX)))
			newI := ((uint64(newX-1)+uint64(newY-1)*size))
			if newI >= 0 && newI < this.Size() {
				n.Write(newI, this.Read(oldI))
			}
			fmt.Printf("%02fx%02f %02d : %02fx%02f %02d\n", oldX, oldY, oldI, newX, newY, newI)
			// fmt.Printf("%02fx%02f %02d\n", oldX, oldY, oldI)
		}
	}
	return n
}
