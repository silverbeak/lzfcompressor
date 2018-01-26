package lzfcompressor

import (
	"errors"

	lzf "github.com/zhuyie/golzf"
)

// CompressLZF Takes payload and returns the compressed data. Please note that the return value also contains the LZF header
func CompressLZF(msg []byte) ([]byte, error) {
	originalLen := len(msg)

	compressed := make([]byte, len(msg))
	compressedCount, err := lzf.Compress(msg, compressed)
	if err != nil {
		return []byte{}, err
	}

	compressed = compressed[:compressedCount]

	header := lzfHeader{
		isCompressed:   true,
		originalLen:    (int64)(originalLen),
		compressedLen:  (int64)(compressedCount),
		compressMethod: "ZV",
	}

	structure := lzfStructure{header: header, body: compressed}

	return structure.toCompressed()
}

// DecompressLZF Takes a compressed LZF payload (with LZF header) and returns the decompressed payload
func DecompressLZF(msg []byte) ([]byte, error) {
	lzfStructure, err := lzfStructureFrom(msg)
	if err != nil {
		return []byte{}, err
	}
	decompressed := make([]byte, lzfStructure.header.originalLen)
	decompressedSize, err := lzf.Decompress(lzfStructure.body, decompressed)
	if err != nil {
		return []byte{}, err
	}
	if decompressedSize == 0 {
		return []byte{}, errors.New("Size of decompressed data must be more than 0")
	}
	return decompressed, nil
}
