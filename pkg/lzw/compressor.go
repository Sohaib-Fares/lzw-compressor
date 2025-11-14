package lzw

import (
	"bytes"
	"errors"
	"fmt"
)

func initialization(ENCODING_BITS int) (int, map[string]int) {

	dictionary := make(map[string]int)

	return 1 << ENCODING_BITS, dictionary
}

func Compress(input []byte, ENCODING_BITS int) ([]byte, error) {

	if len(input) == 0 {
		return nil, errors.New("empty input, please enter a valid input")
	}

	if ENCODING_BITS > 16 || ENCODING_BITS < 9 {
		return nil, errors.New("unsupported encoding, try 9 - 16 bits")
	}

	MAX_DICTIONARY_SIZE, dictionary := initialization(ENCODING_BITS)

	for i := 0; i < 256; i++ {
		dictionary[string(byte(i))] = i
	}

	next_code := 256
	STRING := string(input[0])
	codes := []int{}

	for i := 1; i < len(input); i++ {
		CHAR := string(input[i])
		combined := STRING + CHAR

		if _, exist := dictionary[combined]; exist {
			STRING = combined
		} else {
			codes = append(codes, dictionary[STRING])

			if next_code < MAX_DICTIONARY_SIZE {
				dictionary[combined] = next_code
				next_code++
			}
			STRING = CHAR
		}
	}
	codes = append(codes, dictionary[STRING])
	fmt.Printf("Codes: %v\n", codes) // just for debugging
	return PackCodes(codes, ENCODING_BITS), nil

}

func PackCodes(codes []int, ENCODING_BITS int) []byte {

	var buf bytes.Buffer
	var bitBuffer uint64
	var bitCount uint

	mask := (1 << ENCODING_BITS) - 1

	for _, c := range codes {
		v := uint64(c & mask)
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
