package pefile

import "time"

type ImageFileHeader struct {
	machine              uint16
	numberOfSections     uint16
	timeDateStamp        uint32
	pointerToSymbolTable uint32
	numberOfSymbols      uint32
	sizeOfOptionalHeader int16
	characteristics      int16
}

// for sizeOfOptionalHeader up to array is 96 plus
// the array is 0xEO Should be constant in the world of
// windows but could change

func CreateImageFileHeader(numberOfSections uint16) ImageFileHeader {
	imageFileHeader := ImageFileHeader{
		machine:              0x014c,
		numberOfSections:     numberOfSections,
		timeDateStamp:        uint32(time.Now().Unix()),
		sizeOfOptionalHeader: 0xe0,
		characteristics:      0x0002,
	}

	return imageFileHeader
}
