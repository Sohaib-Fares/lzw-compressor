package bitoperations

import (
	"bytes"
)

// PackCodes packs integer codes into a byte array using the specified number of bits per code
func PackCodes(codes []int, ENCODING_BITS int) []byte {
	var buf bytes.Buffer
	var bitBuffer uint64
	var bitCount uint

	mask := (1 << ENCODING_BITS) - 1

	for _, code := range codes {
		v := uint64(code & mask)
		bitBuffer = (bitBuffer << ENCODING_BITS) | v
		bitCount += uint(ENCODING_BITS)

		for bitCount >= 8 {
			shift := bitCount - 8
			b := byte(bitBuffer >> shift)
			buf.WriteByte(b)
			bitBuffer &= (1 << shift) - 1
			bitCount = shift
		}
	}

	if bitCount > 0 {
		b := byte(bitBuffer << (8 - bitCount))
		buf.WriteByte(b)
	}

	return buf.Bytes()
}
