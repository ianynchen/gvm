package classfile

import "github.com/ianynchen/gvm/jvm/util"

type Interface struct {
	Index uint16
}

func (class *Class) parseInterfaces(content []byte, pos int) (int, error) {
	size, err := util.ParseUint16(content[pos:])
	class.InterfacesCount = size

	if err != nil {
		return pos, err
	}

	pos += 2
	for i := uint16(0); i < class.InterfacesCount; i++ {
		index, e := util.ParseUint16(content[pos:])
		class.Interfaces[i] = Interface{Index: index}

		if e != nil {
			return pos, e
		}
		pos += 2
	}
	return pos, err
}
