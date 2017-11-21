package classfile

import "github.com/ianynchen/gvm/jvm/util"

type Method struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []*Attribute
}

func (class *Class) parseMethod(content []byte, pos int) (int, error) {
	var err error

	class.MethodsCount, err = util.ParseUint16(content[pos:])
	pos += 2
	if err != nil {
		return pos, err
	}

	for i := uint16(0); i < class.MethodsCount; i++ {
		class.Methods[i] = new(Method)
		class.Methods[i].AccessFlags, err = util.ParseUint16(content[pos:])
		pos += 2
		if err != nil {
			return pos, err
		}

		class.Methods[i].NameIndex, err = util.ParseUint16(content[pos:])
		pos += 2
		if err != nil {
			return pos, err
		}

		class.Methods[i].DescriptorIndex, err = util.ParseUint16(content[pos:])
		pos += 2
		if err != nil {
			return pos, err
		}

		class.Methods[i].AttributesCount, err = util.ParseUint16(content[pos:])
		pos += 2
		if err != nil {
			return pos, err
		}
		class.Methods[i].Attributes = make([]*Attribute, int(class.Methods[i].AttributesCount))

		for j := 0; j < int(class.Methods[i].AttributesCount); j++ {
			attribute, p, err := parseAttribute(content, pos)

			if err == nil {
				class.Methods[i].Attributes[j] = attribute
				pos = p
			} else {
				return p, err
			}
		}
	}

	//TODO parse attributes
	return pos, err
}
