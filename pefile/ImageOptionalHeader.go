package pefile

type ImageOptionalHeader struct {
	magic                       uint16            //  2
	majorLinkerVersion          byte              //  3
	minorLinkerVersion          byte              //  4
	sizeOfCode                  uint32            //  8
	sizeOfInitializedData       uint32            // 12
	sizeOfUninitializedData     uint32            // 16
	addressOfEntryPoint         uint32            // 20
	baseOfCode                  uint32            // 24
	baseOfData                  uint32            // 28
	imageBase                   uint32            // 32
	sectionAlignment            uint32            // 36
	fileAlignment               uint32            // 40
	majorOperatingSystemVersion uint16            // 42
	minorOperatingSystemVersion uint16            // 44
	majorImageVersion           uint16            // 46
	minorImageVersion           uint16            // 48
	majorSubsystemVersion       uint16            // 50
	minorSubsystemVersion       uint16            // 52
	reserved1                   uint32            // 56
	sizeOfImage                 uint32            // 60
	sizeOfHeaders               uint32            // 64
	checkSum                    uint32            // 68
	subsystem                   uint16            // 70
	dllCharacteristics          uint16            // 72
	sizeOfStackReserve          uint32            // 76
	sizeOfStackCommit           uint32            // 80
	sizeOfHeapReserve           uint32            // 84
	sizeOfHeapCommit            uint32            // 88
	loaderFlags                 uint32            // 92
	numberOfRvaAndSizes         uint32            // 96
	dataDirectory               [16]DataDirectory // 128
}

type DataDirectory struct {
	virtualAddress uint32
	size           uint32
}

func CreateImageOptionalHeader() ImageOptionalHeader {
	imageOptionalHeader := ImageOptionalHeader{
		magic:                 0x10b,
		addressOfEntryPoint:   0x1000,
		imageBase:             0x400000, // Per MS documenation exe default
		sectionAlignment:      0x1000,
		fileAlignment:         0x0200,
		majorSubsystemVersion: 0x04,
		sizeOfImage:           0x4000,
		sizeOfHeaders:         0x0200,
		subsystem:             0x03,
		numberOfRvaAndSizes:   0x10,
	}

	imageOptionalHeader.dataDirectory[1].virtualAddress = 0x2000

	return imageOptionalHeader
}
