package lzfcompressor

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

type lzfHeader struct {
	compressMethod string
	isCompressed   bool
	originalLen    int64
	compressedLen  int64
}

type lzfStructure struct {
	header lzfHeader
	body   []byte
}

func lzfStructureFrom(barray []byte) (lzfStructure, error) {
	header, err := headerFromCompressed(barray[:7])
	if err != nil {
		return lzfStructure{}, err
	}
	var body []byte

	if header.isCompressed {
		body = barray[7:]
	} else {
		body = barray[5:]
	}

	return lzfStructure{header: header, body: body}, nil
}

func headerFromCompressed(header []byte) (lzfHeader, error) {
	isCompressed := header[2] != 0
	var originalLen int64
	var err error

	if isCompressed {
		str := fmt.Sprintf("%x", header[5:7])
		originalLen, err = strconv.ParseInt(str, 16, 32)
		if err != nil {
			return lzfHeader{originalLen: originalLen}, err
		}
	}

	str := fmt.Sprintf("%x", header[3:5])
	compressedLen, err := strconv.ParseInt(str, 16, 32)
	if err != nil {
		return lzfHeader{}, err
	}

	return lzfHeader{isCompressed: isCompressed, originalLen: originalLen, compressedLen: compressedLen}, nil
}

func itoabarray(i int64) ([]byte, error) {
	o := new(bytes.Buffer)
	err := binary.Write(o, binary.BigEndian, i)
	if err != nil {
		return []byte{}, err
	}
	byt := o.Bytes()[6:]
	return byt, nil
}

func (structure lzfStructure) toCompressed() ([]byte, error) {

	header := structure.header
	var isCompressed byte
	if header.isCompressed {
		isCompressed = 0x01
	} else {
		isCompressed = 0x00
	}

	compressedLen, _ := itoabarray(header.compressedLen)

	originalLen, _ := itoabarray(header.originalLen)

	var a1 []byte
	a1 = append(
		[]byte(header.compressMethod),
		isCompressed)

	a2 := append(
		a1,
		compressedLen...,
	)

	a3 := append(
		a2,
		originalLen...,
	)

	a4 := append(
		a3,
		structure.body...,
	)
	return a4, nil
}
