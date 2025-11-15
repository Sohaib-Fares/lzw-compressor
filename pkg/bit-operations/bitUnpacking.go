package bitoperations

// UnpackCodes unpacks codes from a byte array using the specified number of bits per code
func UnpackCodes(input []byte, ENCODING_BITS int) []int {
	var codes []int
	var bitBuffer uint64
	var bitCount uint

	mask := uint64((1 << ENCODING_BITS) - 1)

	for _, b := range input {
		bitBuffer = (bitBuffer << 8) | uint64(b)
		bitCount += 8

		for bitCount >= uint(ENCODING_BITS) {
			shift := bitCount - uint(ENCODING_BITS)
			code := int((bitBuffer >> shift) & mask)
			codes = append(codes, code)
			bitBuffer &= (1 << shift) - 1
			bitCount = shift
		}
	}

	return codes
}
