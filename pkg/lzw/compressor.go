package lzw

import (
	"bytes"
	"fmt"
)

const (
	ENCODING_BITS       = 12
	MAX_DICTIONARY_SIZE = 4096
)

func Compress(input []byte) []byte {

	dictionary := make(map[string]int)

	for i := range 256 {
		dictionary[string(byte(i))] = i
	}

	next_code := 256

	STRING := string(input[0])
	codes := []int{}
	// TODO: change int implementation to a more proper type management
	// TODO: add support for other bit encodings and handle full map edge cases
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
	return PackCodes(codes)

}

func PackCodes(codes []int) []byte {
	var buf bytes.Buffer

	for i := 0; i < len(codes); i += 2 {
		if i+1 < len(codes) {
			// case we have 2 bytes left
			code1 := codes[i]
			code2 := codes[i+1]

			buf.WriteByte(byte(code1 >> 4))
			buf.WriteByte(byte((code1&0x0F)<<4 | (code2>>8)&0x0F))
			buf.WriteByte(byte(code2 & 0xFF))

		} else {
			// case we're in the last byte
			code1 := codes[i]
			buf.WriteByte(byte(code1 >> 4))
			buf.WriteByte(byte((code1 & 0x0F) << 4))

		}
	}

	return buf.Bytes()
}
