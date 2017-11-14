package classfile

import (
	"errors"

	"github.com/ianynchen/gvm/jvm/util"
)

const (
	CONSTANTClass              uint8 = 7
	CONSTANTFieldref           uint8 = 9
	CONSTANTMethodref          uint8 = 10
	CONSTANTInterfaceMethodref uint8 = 11
	CONSTANTString             uint8 = 8
	CONSTANTInteger            uint8 = 3
	CONSTANTFloat              uint8 = 4
	CONSTANTLong               uint8 = 5
	CONSTANTDouble             uint8 = 6
	CONSTANTNameAndType        uint8 = 12
	CONSTANTUtf8               uint8 = 1
	CONSTANTMethodHandle       uint8 = 15
	CONSTANTMethodType         uint8 = 16
	CONSTANTInvokeDynamic      uint8 = 18
)

/*
ConstantPoolInfo contains the information for each
constant pool item. The first field is a uint8 tag that
indicates the type of the actual constant pool item contained, and
the second field is an interface{} pointing to the actual constant
pool item.
*/
type ConstantPoolInfo struct {
	Tag  uint8
	Info interface{}
}

/*
ConstantClassInfo contains the index into a constant pool info
that is of UTF8 type, the NameIndex is the position of the item
in the constant pool, and the value of that UTF8 string is the name
of the class.
*/
type ConstantClassInfo struct {
	NameIndex uint16
}

/*
ConstantReference serves as the base struct for ConstantMethodref,
ConstantFieldref and ConstantInterfaceMethodref fields. It contains
two fields that are used to identify the position of the actual item.
The ClassIndex and the NameAndTypeIndex
*/
type ConstantReference struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantMethodRef struct {
	ConstantReference
}

type ConstantFieldref struct {
	ConstantReference
}

type ConstantInterfaceMethodref struct {
	ConstantReference
}

type ConstantStringInfo struct {
	Index uint16
}

type ConstantUtf8Info struct {
	Length uint8
	Info   string
}

func (class *Class) Name(index uint16) (string, error) {

	if index >= 1 && int(index) < len(class.ConstantPool) {
		if class.ConstantPool[index].Tag == CONSTANTUtf8 {
			if value, ok := class.ConstantPool[index].Info.(string); ok {
				return value, nil
			}
			return "", errors.New("Invalid string for name")
		}
	}
	return "", errors.New("Invalid position index for name")
}

func (class *Class) String(index uint16) (string, error) {
	return class.Name(index)
}

func parseConstantClassInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantClassInfo)
	index, err := util.ParseUint16(content)
	result.NameIndex = index
	return result, 2, err
}
