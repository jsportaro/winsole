package pefile

type DosHeader struct {
	signature [2]byte
	lastsize  int16
	nblocks   int16
	nreloc    int16
	hdrsize   int16
	minalloc  int16
	maxalloc  int16
	ss        int16
	sp        int16
	checksum  int16
	ip        int16
	cs        int16
	relocpos  int16
	noverlay  int16
	reserved1 [4]int16
	oem_id    int16
	oem_info  int16
	reserved2 [10]int16
	e_lfanew  int32
}

func CreateDosHeader() DosHeader {
	dosHeader := DosHeader{
		signature: [2]byte{'M', 'Z'},
	}
	return dosHeader
}
