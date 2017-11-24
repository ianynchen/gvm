package classfile

import (
	"github.com/ianynchen/gvm/jvm/util"
)

type Interface struct {
	Index uint16 `json:"index"`
}

func (class *Class) parseInterfaces(content []byte, pos int) (int, error) {
	size, err := util.ParseUint16(content[pos:])
	class.InterfacesCount = size
	pos += 2

	if err != nil {
		return pos, err
	}

	class.Interfaces = make([]*Interface, size)
	for i := uint16(0); i < class.InterfacesCount; i++ {
		index, e := util.ParseUint16(content[pos:])
		class.Interfaces[i] = new(Interface)
		class.Interfaces[i].Index = index

		if e != nil {
			return pos, e
		}
		pos += 2
	}
	return pos, err
}
