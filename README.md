# LZW Compressor

A fast and efficient implementation of the Lempel-Ziv-Welch (LZW) compression algorithm in Go.

## Description

LZW Compressor is a command-line tool that implements the LZW lossless data compression algorithm. The tool supports both compression and decompression of files using configurable bit encoding (9-16 bits).

### Key Features

- **Compression**: Reduce file sizes by identifying and replacing repeated patterns with shorter codes
- **Decompression**: Restore compressed files to their original form without any data loss
- **Efficiency**: Particularly effective for files with repetitive patterns (text files, logs, etc.)
- **Flexibility**: Supports multiple encoding bit sizes (9-16 bits) to optimize compression ratios

## Installation

### Prerequisites

- Go 1.16 or higher

### Install from source

```bash
# Clone the repository
git clone https://github.com/Sohaib-Fares/lzw-compressor.git
cd lzw-compressor

# Build the binary
go build -o lzw-compressor ./cmd/lzw-compressor

# Optionally, install to your Go bin directory
go install ./cmd/lzw-compressor
```

### Verify installation

```bash
./lzw-compressor -h
```

## Usage

### Basic Compression

Compress a file using default settings (12-bit encoding):

```bash
./lzw-compressor -f input.txt
```

This creates `input.txt.lzw` in the same directory.

### Basic Decompression

Decompress a `.lzw` file:

```bash
./lzw-compressor -f input.txt.lzw -d
```

This creates `input.txt` (original filename without `.lzw` extension).

### Advanced Options

#### Custom output file

```bash
# Compression with custom output
./lzw-compressor -f input.txt -o compressed.dat

# Decompression with custom output
./lzw-compressor -f compressed.dat -d -o output.txt
```

#### Custom bit encoding

Use different bit sizes (9-16) for compression:

```bash
# Compress with 9 bits (smaller dictionary, less compression)
./lzw-compressor -f input.txt -b 9

# Compress with 16 bits (larger dictionary, better compression for large files)
./lzw-compressor -f input.txt -b 16
```

**Note**: Use the same bit size for both compression and decompression.

### Command-Line Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-f` | Input file path (required) | - |
| `-o` | Output file path | `<input>.lzw` for compression<br>`<input without .lzw>` for decompression |
| `-b` | Number of encoding bits (9-16) | 12 |
| `-d` | Enable decompression mode | false (compression mode) |

### Examples

```bash
# Compress a text file
./lzw-compressor -f document.txt

# Compress with 14-bit encoding
./lzw-compressor -f largefile.log -b 14

# Decompress
./lzw-compressor -f document.txt.lzw -d

# Compress with custom output
./lzw-compressor -f data.bin -o data.compressed -b 16

# Decompress with custom output and matching bit size
./lzw-compressor -f data.compressed -d -o data.original -b 16
```

### Output Information

The tool displays compression statistics:

```
Compression successful!
Input size:  134 bytes
Output size: 131 bytes
Ratio:       97.76%
Output written to: document.txt.lzw
```

## Project Structure

```
lzw-compressor/
├── cmd/
│   └── lzw-compressor/
│       └── main.go              # CLI application entry point
├── pkg/
│   ├── bit-operations/
│   │   ├── bitPacking.go        # Pack codes into bytes
│   │   ├── bitUnpacking.go      # Unpack bytes into codes
│   │   └── bitOperations_test.go # Bit operation tests
│   └── lzw/
│       ├── compressor.go        # LZW compression algorithm
│       ├── decompressor.go      # LZW decompression algorithm
│       └── compressor_test.go   # LZW algorithm tests
├── go.mod
└── README.md
```

## Performance

The LZW algorithm works best with:
-  Text files with repetitive patterns
-  Log files
-  Source code files
-  Configuration files

Less effective with:
-  Already compressed files (ZIP, GZIP, etc.)
-  Random or encrypted data
-  Small files (compression overhead)

### Choosing Bit Size

- **9 bits**: Faster, smaller dictionary (512 codes), lower compression ratio
- **12 bits**: Balanced performance (4096 codes) - **recommended default**
- **16 bits**: Slower, larger dictionary (65536 codes), better for large files

## Algorithm Details

LZW (Lempel-Ziv-Welch) is a universal lossless data compression algorithm that:

1. Builds a dictionary of input sequences during compression
2. Replaces repeated sequences with dictionary codes
3. Adapts to the input data without prior knowledge
4. Reconstructs the dictionary during decompression

This implementation uses:
- Dynamic dictionary building (256 initial entries for single bytes)
- Configurable code bit-width (9-16 bits)
- Efficient bit packing for compressed output

## License

MIT License

