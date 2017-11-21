package classfile

import (
	"errors"

	"github.com/ianynchen/gvm/jvm/util"
)

type AttributeBase struct {
	AttributeNameIndex uint16 `json:"attribute_name_index"`
	AttributeLength    uint32 `json:"attribute_length"`
}
type Attribute struct {
	AttributeBase
	Content []byte
}

func parseAttribute(content []byte, pos int) (*Attribute, int, error) {
	nameIndex, err1 := util.ParseUint16(content[pos:])

	if err1 != nil {
		return nil, pos + 2, err1
	}
	pos += 2

	length, err2 := util.ParseUint32(content[pos:])
	if err2 != nil {
		return nil, pos + 4, err2
	}
	pos += 4

	attribute := new(Attribute)
	attribute.AttributeNameIndex = nameIndex
	attribute.AttributeLength = length
	attribute.Content = make([]byte, int(length))
	copied := copy(attribute.Content, content[pos:])

	if uint32(copied) != length {
		return attribute, pos + copied, errors.New("Unexpected length of array for attribute")
	}
	return attribute, pos + int(length), nil
}

type ConstantValueAttribute struct {
	AttributeBase
	ConstantValueIndex uint16 `json:"constant_value_index"`
}

type ExceptionTableEntry struct {
	StartPC   uint16
	EndPC     uint16
	HandlerPC uint16
	CatchType uint16
}

type CodeAttribute struct {
	AttributeBase
	MaximumStack          uint16
	MaximumLocals         uint16
	CodeLength            uint32
	Code                  []byte
	ExceptionTableLength  uint16
	ExceptionTableEntries []ExceptionTableEntry
	AttributesCount       uint16
	AttributeInfo         []interface{}
}
