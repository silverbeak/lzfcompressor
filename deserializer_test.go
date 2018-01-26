package lzfcompressor

import (
	"testing"
)

func getCompressed(t *testing.T) lzfStructure {
	content := []byte(ipsum)

	compressedStruct, err := lzfStructureFrom(content)
	if err != nil {
		t.Fatalf("Error when compressing: %v", err)
	}
	return compressedStruct
}

func TestCompression(t *testing.T) {
	content := []byte(ipsum)

	compressed, err := CompressLZF(content)
	if err != nil {
		t.Fatalf("Error when compressing: %v", err)
	}

	compressedStruct, err := lzfStructureFrom(compressed)
	if err != nil {
		t.Fatalf("Could not create compressed struct: %v", err)
	}

	if compressedStruct.header.originalLen != 970 {
		t.Fatalf("Original length did not match. Should have been 970, was %v", compressedStruct.header.originalLen)
	}

	if compressedStruct.header.compressedLen != 746 {
		t.Fatalf("Original length did not match. Should have been 746, was %v", compressedStruct.header.compressedLen)
	}
}

func TestLzfStructurePackaging(t *testing.T) {
	content := []byte(ipsum)

	compressed, err := CompressLZF(content)
	if err != nil {
		t.Fatalf("Error when compressing: %v", err)
	}

	decompressed, err := DecompressLZF(compressed)
	if err != nil {
		t.Fatalf("Could not decompress lzf %v", err)
	}

	if len(decompressed) != len(content) {
		t.Fatalf("Length did not match for decompressed stuff. Wanted %v, got %v", len(content), len(decompressed))
	}
}

var ipsum = `But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness. No one rejects, dislikes, or avoids pleasure itself, because it is pleasure, but because those who do not know how to pursue pleasure rationally encounter consequences that are extremely painful. Nor again is there anyone who loves or pursues or desires to obtain pain of itself, because it is pain, but because occasionally circumstances occur in which toil and pain can procure him some great pleasure. To take a trivial example, which of us ever undertakes laborious physical exercise, except to obtain some advantage from it? But who has any right to find fault with a man who chooses to enjoy a pleasure that has no annoying consequences, or one who avoids a pain that produces no resultant pleasure?`
