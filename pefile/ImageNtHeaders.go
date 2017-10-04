package pefile

type ImageNtHeaders struct {
	signature      [2]byte
	padding        int16
	fileHeader     ImageFileHeader
	optionalHeader ImageOptionalHeader
}

func CreateImageNtHeaders(numberOfSections uint16) ImageNtHeaders {
	imageNtHeaders := ImageNtHeaders{
		signature:      [2]byte{'P', 'E'},
		fileHeader:     CreateImageFileHeader(numberOfSections),
		optionalHeader: CreateImageOptionalHeader(),
	}

	return imageNtHeaders
}
