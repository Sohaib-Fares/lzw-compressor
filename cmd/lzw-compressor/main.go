package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/Sohaib-Fares/lzw-compressor/pkg/lzw"
)

func main() {
	fileptr := flag.String("f", "", "Specified file to compress/decompress (required)")
	outptr := flag.String("o", "", "Output file path (defaults to <input>.lzw for compression or <input>.decompressed for decompression)")
	ENCODING_BITS := flag.Int("b", 12, "Number of bits used to encode compressed file (9-16)")
	decompress := flag.Bool("d", false, "Decompress mode (default is compress)")
	flag.Parse()

	if *fileptr == "" {
		flag.Usage()
		log.Fatal("error: a file path is required. Use -f <path>")
	}

	input, err := os.ReadFile(*fileptr)
	if err != nil {
		log.Fatalf("error while reading file %q: %v", *fileptr, err)
	}
	if len(input) == 0 {
		log.Fatal("error: empty file")
	}

	var output []byte
	var defaultOutput string

	if *decompress {

		output, err = lzw.Decompress(input, *ENCODING_BITS)
		if err != nil {
			log.Fatalf("decompression failed: %v", err)
		}

		if strings.HasSuffix(*fileptr, ".lzw") {
			defaultOutput = strings.TrimSuffix(*fileptr, ".lzw")
		} else {
			defaultOutput = fmt.Sprintf("%s.decompressed", *fileptr)
		}

		fmt.Printf("Decompression successful!\n")
		fmt.Printf("Input size:  %d bytes\n", len(input))
		fmt.Printf("Output size: %d bytes\n", len(output))
	} else {
		output, err = lzw.Compress(input, *ENCODING_BITS)
		if err != nil {
			log.Fatalf("compression failed: %v", err)
		}

		defaultOutput = fmt.Sprintf("%s.lzw", *fileptr)

		compressionRatio := float64(len(output)) / float64(len(input)) * 100
		fmt.Printf("Compression successful!\n")
		fmt.Printf("Input size:  %d bytes\n", len(input))
		fmt.Printf("Output size: %d bytes\n", len(output))
		fmt.Printf("Ratio:       %.2f%%\n", compressionRatio)
	}

	if *outptr == "" {
		*outptr = defaultOutput
	}

	if err := os.WriteFile(*outptr, output, 0644); err != nil {
		log.Fatalf("error writing output file %q: %v", *outptr, err)
	}

	fmt.Printf("Output written to: %s\n", *outptr)
}
