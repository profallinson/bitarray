package bitarray

import (
	"math"
)

func (this *bitArray) Radius(x, y, r int) BitArray {
	max := int(math.Sqrt(float64(this.Size())))
	xStart := x - r
	yStart := y - r
	xEnd := x + r
	yEnd := y + r
	if xStart < 0 {
		xStart = 0
	}
	if yStart < 0 {
		yStart = 0
	}
	if xEnd > max {
		xEnd = max
	}
	if yEnd > max {
		yEnd = max
	}
	ret := make([]bool, 0)
	for ; yStart < yEnd; yStart++ {
		for i := xStart; i < xEnd; i++ {
			ret = append(ret, this.Read(uint64(yStart*max+i)))
		}
	}
	return NewBitArrayFromBools(ret)
}

/*
	These are old ideas but here just in case.
*/

// If the bit array length can be squared to an `uint64` then it's content are rotated by the given `deg`.
func (this *bitArray) Rotate(deg float64) BitArray {
	sqrt := math.Ceil(math.Sqrt(float64(this.Size())))
	offset := sqrt/2 - 0.5
	n := NewBitArrayOfLength(this.Size())
	sin, cos := math.Sincos(deg * math.Pi / 180)
	oldI := uint64(0)
	for oldY := offset * -1; oldY <= offset; oldY++ {
		for oldX := offset * -1; oldX <= offset; oldX++ {
			newX := math.Abs(math.Round(((oldX * cos) - (oldY * sin)) + offset)) // x’ = Xcos - Ysin
			newY := math.Abs(math.Round(((oldX * sin) + (oldY * cos)) + offset)) // y’ = Xsin + Ycos
			newI := uint64((newY * sqrt) + newX)
			if oldI >= 0 && oldI < this.Size() && newI >= 0 && newI < this.Size() {
				n.Write(newI, this.Read(oldI))
			}
			oldI++
		}
	}
	return n
}

func (this *bitArray) Move(x int, y int) BitArray {
	sqrt := math.Ceil(math.Sqrt(float64(this.Size())))
	size := int(sqrt)
	n := NewBitArrayOfLength(this.Size())
	oldI := uint64(0)
	for newY := 0 - y; newY < size-y; newY++ {
		for newX := 0 + x; newX < size+x; newX++ {
			if newX >= 0 && newX < size && newY >= 0 && newY < size {
				newI := uint64((newY * size) + newX)
				if oldI >= 0 && oldI < this.Size() && newI >= 0 && newI < this.Size() {
					n.Write(newI, this.Read(oldI))
				}
			}
			oldI++
		}
	}
	return n
}

func (this *bitArray) Clip(low uint8, high uint8, byteGroupSize uint64) BitArray {
	source := this.ToBytes()
	bytes := make([]byte, len(source))
	for g := uint64(0); g <= uint64(len(source))-byteGroupSize; g += byteGroupSize {
		test := make([]byte, 0)
		for i := uint64(0); i < byteGroupSize; i++ {
			if source[g+i] >= low && source[g+i] <= high {
				test = append(test, source[g+i])
			}
		}
		if uint64(len(test)) == byteGroupSize {
			for i, v := range test {
				bytes[g+uint64(i)] = v
			}
		}
	}
	return NewBitArrayFromBytes(bytes)
}

func (this *bitArray) Contrast(low uint8, high uint8, byteGroupSize uint64) BitArray {
	source := this.ToBytes()
	bytes := make([]byte, len(source))
	for g := uint64(0); g <= uint64(len(source))-byteGroupSize; g += byteGroupSize {
		test := make([]byte, 0)
		for i := uint64(0); i < byteGroupSize; i++ {
			if source[g+i] >= low && source[g+i] <= high {
				test = append(test, source[g+i])
			}
		}
		if uint64(len(test)) == byteGroupSize {
			for i, _ := range test {
				bytes[g+uint64(i)] = byte(255)
			}
		}
	}
	return NewBitArrayFromBytes(bytes)
}

func (this *bitArray) Avg(byteGroupSize uint64) BitArray {
	source := this.ToBytes()
	bytes := make([]byte, int(math.Ceil((float64(uint64(len(source)) / byteGroupSize)))))
	for i := uint64(0); i < uint64(len(source)); i += byteGroupSize {
		byteGroup := source[i : i+byteGroupSize]
		sum := 0
		for _, v := range byteGroup {
			sum += int(v)
		}
		bytes[i/byteGroupSize] = byte(sum / len(byteGroup))
	}
	return NewBitArrayFromBytes(bytes)
}

func (this *bitArray) Binarify(low uint8, high uint8) BitArray {
	source := this.Contrast(low, high, 1).ToBytes()
	bitArray := NewBitArrayOfLength(uint64(len(source)))
	for i, v := range source {
		bitArray.Write(uint64(i), v >= low && v <= high)
	}
	return bitArray
}
