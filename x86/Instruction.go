package x86

import (
	"bytes"
)

//Instruction - Base for all instructions
type Instruction interface {
	Write(b bytes.Buffer)
}
