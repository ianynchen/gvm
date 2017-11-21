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

	for i := uint16(0); i < class.FieldsCount; i++ {
		class.Fields[i].AccessFlags, err = util.ParseUint16(content[pos:])
		pos += 2
		if err != nil {
			return pos, err
		}

		class.Fields[i].NameIndex, err = util.ParseUint16(content[pos:])
		pos += 2
		if err != nil {
			return pos, err
		}

		class.Fields[i].DescriptorIndex, err = util.ParseUint16(content[pos:])
		pos += 2
		if err != nil {
			return pos, err
		}

		class.Fields[i].AttributesCount, err = util.ParseUint16(content[pos:])
		pos += 2
		if err != nil {
			return pos, err
		}
		class.Fields[i].Attributes = make([]*Attribute, int(class.Fields[i].AttributesCount))

		//TODO add attribute parsing
		for j := 0; j < int(class.Fields[i].AttributesCount); j++ {
			attribute, p, err := parseAttribute(content, pos)

			if err == nil {
				class.Fields[i].Attributes[j] = attribute
				pos = p
			} else {
				return p, err
			}
		}
	}
	return pos, err
}
