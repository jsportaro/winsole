package x86

import (
	"bytes"
	"encoding/binary"
)

//ObjectBuilder - For now, an EXE
type ObjectBuilder struct {
	executableBytes bytes.Buffer
}

//NewExecutable - Creates a new *ObjectBuilder
func NewExecutable() *ObjectBuilder {

	objectBuilder := &ObjectBuilder{}

	return objectBuilder
}

//AddI32 Add add 32 bit immediate to EAX
func (o *ObjectBuilder) AddI32(constant int32) *ObjectBuilder {

	binary.Write(&o.executableBytes, binary.LittleEndian, uint8(0x05))
	binary.Write(&o.executableBytes, binary.LittleEndian, constant)

	return o
}

//MoveI32 Move 32 bit immediate to EAX
func (o *ObjectBuilder) MoveI32(constant int32) *ObjectBuilder {

	binary.Write(&o.executableBytes, binary.LittleEndian, uint8(0xB8))
	binary.Write(&o.executableBytes, binary.LittleEndian, uint32(constant))

	return o
}

//PushEax
func (o *ObjectBuilder) PushEax() *ObjectBuilder {

	binary.Write(&o.executableBytes, binary.LittleEndian, uint8(0x50))

	return o
}

//GetBytes return them
func (o *ObjectBuilder) GetBytes() []byte {

	binary.Write(&o.executableBytes, binary.LittleEndian, uint16(0x15FF))
	binary.Write(&o.executableBytes, binary.LittleEndian, uint32(0x402068))

	return o.executableBytes.Bytes()
}

func (o *ObjectBuilder) writeOpCode(opcode int32) {
	binary.Write(&o.executableBytes, binary.LittleEndian, opcode)
}
