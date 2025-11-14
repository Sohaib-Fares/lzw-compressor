package main

import (
	"fmt"

	"github.com/Sohaib-Fares/lzw-compressor/pkg/lzw"
)

//TODO: Implement CLI logic using flag lib.
//TODO: Remove test code and implement proper error handlig logic

func main() {
	testStrings := []string{
		"ABABABA",
		"ABCABCABCABC",
		"ABABABCABABABCABABABCABABABCXYZXYZXYZ",
	}

	for _, test := range testStrings {
		input := []byte(test)
		compressed := lzw.Compress(input)

		ratio := float64(len(input)) / float64(len(compressed))
		fmt.Printf("Input: %s\n", test)
		fmt.Printf("Original size: %d bytes\n", len(input))
		fmt.Printf("Compressed size: %d bytes\n", len(compressed))
		fmt.Printf("Compression ratio: %.2f:1\n", ratio)
		fmt.Println("output: ", compressed)
	}
}
