package pefile

type ImageSectionHeader struct {
	name                 [8]byte
	address              uint32
	virtualAddress       uint32
	sizeOfRawData        uint32
	pointerToRawData     uint32
	pointerToRelocations uint32
	pointerToLineNumbers uint32
	numberOfRelocations  uint16
	numberOfLineNumbers  uint16
	charactaristics      uint32
}
