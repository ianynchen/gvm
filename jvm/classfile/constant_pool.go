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

type ConstantMethodref struct {
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

type ConstantIntegerInfo struct {
	Value int32
}

type ConstantFloadInfo struct {
	Value util.Numerical
}

type ConstantLongInfo struct {
	Value int64
}

type ConstantDoubleInfo struct {
	Value util.Numerical
}

type ConstantUtf8Info struct {
	Length uint16
	Info   string
}

type ConstantNameAndTypeInfo struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

type ConstantMethodHandleInfo struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

type ConstantMethodTypeInfo struct {
	DescriptorIndex uint16
}

type ConstantInvokeDynamicInfo struct {
	BootstrapMethodAttributeIndex uint16
	NameAndTypeIndex              uint16
}

func (class *Class) Name(index uint16) (string, error) {

	if index >= 1 && int(index) < len(class.ConstantPool) && class.ConstantPool[index] != nil {
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

func parseConstantFieldrefInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantFieldref)
	index, err := util.ParseUint16(content)

	if err != nil {
		return result, 2, err
	}
	result.ClassIndex = index
	index, err = util.ParseUint16(content[2:])
	result.NameAndTypeIndex = index
	return result, 4, err
}

func parseConstantMethodrefInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantMethodref)
	index, err := util.ParseUint16(content)

	if err != nil {
		return result, 2, err
	}
	result.ClassIndex = index
	index, err = util.ParseUint16(content[2:])
	result.NameAndTypeIndex = index
	return result, 4, err
}

func parseConstantInterfaceMethodrefInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantInterfaceMethodref)
	index, err := util.ParseUint16(content)

	if err != nil {
		return result, 2, err
	}
	result.ClassIndex = index
	index, err = util.ParseUint16(content[2:])
	result.NameAndTypeIndex = index
	return result, 4, err
}

func parseConstantStringInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantStringInfo)
	index, err := util.ParseUint16(content)
	result.Index = index
	return result, 2, err
}

func parseConstantIntegerInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantIntegerInfo)
	value, err := util.ParseUint32(content)
	result.Value = int32(value)
	return result, 4, err
}

func parseConstantFloatInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantFloadInfo)
	value, err := util.ParseFloat32(content)
	result.Value = value
	return result, 4, err
}

func parseConstantLongInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantLongInfo)
	value, err := util.ParseUint64(content)
	result.Value = int64(value)
	return result, 8, err
}

func parseConstantDoubleInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantDoubleInfo)
	value, err := util.ParseFloat64(content)
	result.Value = value
	return result, 8, err
}

func parseConstantNameAndTypeInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantNameAndTypeInfo)
	index, err := util.ParseUint16(content)
	result.NameIndex = index

	if err != nil {
		return result, 2, err
	}
	index, err = util.ParseUint16(content[2:])
	result.DescriptorIndex = index
	return result, 4, err
}

func parseConstantUtf8Info(content []byte) (interface{}, int, error) {
	result := new(ConstantUtf8Info)
	length, err := util.ParseUint16(content)
	result.Length = length

	if err != nil {
		return result, 2, err
	}

	l := 2 + int(length)
	if len(content) < l {
		return result, 2, errors.New("not long enough to parse utf8 string")
	}
	result.Info = string(content[2:l])
	return result, l, err
}

func parseConstantMethodHandleInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantMethodHandleInfo)
	kind, err := util.ParseUint8(content)
	result.ReferenceKind = kind

	if err != nil {
		return result, 1, err
	}

	index, err := util.ParseUint16(content[1:])
	result.ReferenceIndex = index
	return result, 1 + 2, err
}

func parseConstantMethodTypeInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantMethodTypeInfo)
	index, err := util.ParseUint16(content)
	result.DescriptorIndex = index
	return result, 2, err
}

func parseConstantInvokeDynamicInfo(content []byte) (interface{}, int, error) {
	result := new(ConstantInvokeDynamicInfo)
	index, err := util.ParseUint16(content)
	result.BootstrapMethodAttributeIndex = index

	if err != nil {
		return result, 2, err
	}
	index, err = util.ParseUint16(content[2:])
	result.NameAndTypeIndex = index
	return result, 4, err
}
