package lzw

import (
	"errors"
	bitoperations "github.com/Sohaib-Fares/lzw-compressor/pkg/bit-operations"
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
		dictionary[string([]byte{byte(i)})] = i
	}

	next_code := 256
	STRING := string([]byte{input[0]})
	codes := []int{}

	for i := 1; i < len(input); i++ {
		CHAR := string([]byte{input[i]})
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
	return bitoperations.PackCodes(codes, ENCODING_BITS), nil

}
