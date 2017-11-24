package classfile

import "github.com/ianynchen/gvm/jvm/util"

type Field struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []*Attribute
}

func (class *Class) parseFieldInfo(content []byte, pos int) (int, error) {
	size, err := util.ParseUint16(content[pos:])
	class.FieldsCount = size
	pos += 2
	if err != nil {
		return pos, err
	}

	class.Fields = make([]*Field, size)
	for i := uint16(0); i < class.FieldsCount; i++ {
		class.Fields[i] = new(Field)
		offset, err1 := class.Fields[i].parseField(content, pos)

		if err1 == nil {
			return offset, err1
		}
		pos = offset
	}
	return pos, err
}

func (field *Field) parseAccessFlags(content []byte, pos int) (int, error) {
	var err error
	field.AccessFlags, err = util.ParseUint16(content[pos:])
	pos += 2
	return pos, err
}

func (field *Field) parseNameIndex(content []byte, pos int) (int, error) {
	var err error
	field.NameIndex, err = util.ParseUint16(content[pos:])
	pos += 2
	return pos, err
}

func (field *Field) parseDescriptorIndex(content []byte, pos int) (int, error) {
	var err error
	field.DescriptorIndex, err = util.ParseUint16(content[pos:])
	pos += 2
	return pos, err
}

func (field *Field) parseAttributeCount(content []byte, pos int) (int, error) {
	var err error
	field.AttributesCount, err = util.ParseUint16(content[pos:])
	pos += 2
	return pos, err
}

func (field *Field) parseAttributes(content []byte, pos int) (int, error) {
	var err error
	var attribute *Attribute
	field.Attributes = make([]*Attribute, int(field.AttributesCount))

	for j := 0; j < int(field.AttributesCount); j++ {
		attribute, pos, err = parseAttribute(content, pos)

		if err == nil {
			field.Attributes[j] = attribute
		} else {
			return pos, err
		}
	}
	return pos, err
}

func (field *Field) parseField(content []byte, pos int) (int, error) {
	context := parsingContext{content: content, offset: pos, err: nil}
	context.parse(field.parseAccessFlags)
	context.parse(field.parseNameIndex)
	context.parse(field.parseDescriptorIndex)
	context.parse(field.parseAttributeCount)
	context.parse(field.parseAttributes)
	return context.offset, context.err
}
