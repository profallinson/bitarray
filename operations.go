package bitarray

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
	var i uint64
	for i = 0; i < this.Size(); i++ {
		if this.Read(i) {
			n++
		}
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
	var i uint64
	for i = 0; i < this.Size(); i++ {
		if this.Read(i) == a.Read(i) {
			o++
		}
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
func (this *bitArray) Sparsity() float32 {
	return float32(this.Size()) / 100 * float32(this.Norm())
}
