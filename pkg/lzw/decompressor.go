package lzw

import (
	"bytes"
	"errors"
	bitoperations "github.com/Sohaib-Fares/lzw-compressor/pkg/bit-operations"
)

func Decompress(input []byte, ENCODING_BITS int) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("empty input, please enter a valid compressed data")
	}

	if ENCODING_BITS > 16 || ENCODING_BITS < 9 {
		return nil, errors.New("unsupported encoding, try 9 - 16 bits")
	}

	codes := bitoperations.UnpackCodes(input, ENCODING_BITS)
	if len(codes) == 0 {
		return nil, errors.New("no codes found in input")
	}

	MAX_DICTIONARY_SIZE := 1 << ENCODING_BITS
	dictionary := make(map[int][]byte)

	for i := 0; i < 256; i++ {
		dictionary[i] = []byte{byte(i)}
	}

	next_code := 256
	var output bytes.Buffer

	oldCode := codes[0]
	if _, exist := dictionary[oldCode]; !exist {
		return nil, errors.New("invalid compressed data: first code not in dictionary")
	}

	STRING := dictionary[oldCode]
	output.Write(STRING)

	for i := 1; i < len(codes); i++ {
		code := codes[i]
		var entry []byte

		if val, exist := dictionary[code]; exist {
			entry = val
		} else if code == next_code {

			entry = make([]byte, len(STRING)+1)
			copy(entry, STRING)
			entry[len(STRING)] = STRING[0]
		} else {
			return nil, errors.New("invalid compressed data: code not in dictionary")
		}

		output.Write(entry)

		if next_code < MAX_DICTIONARY_SIZE {
			newEntry := make([]byte, len(STRING)+1)
			copy(newEntry, STRING)
			newEntry[len(STRING)] = entry[0]
			dictionary[next_code] = newEntry
			next_code++
		}

		STRING = entry
	}

	return output.Bytes(), nil
}
