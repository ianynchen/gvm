package classfile

const (
	ACC_PUBLIC     uint16 = 0x0001 //Declared public; may be accessed from outside its package.
	ACC_FINAL      uint16 = 0x0010 //Declared final; no subclasses allowed.
	ACC_SUPER      uint16 = 0x0020 //Treat superclass methods specially when invoked by the invokespecial instruction.
	ACC_INTERFACE  uint16 = 0x0200 //Is an interface, not a class.
	ACC_ABSTRACT   uint16 = 0x0400 //Declared abstract; must not be instantiated.
	ACC_SYNTHETIC  uint16 = 0x1000 //Declared synthetic; not present in the source code.
	ACC_ANNOTATION uint16 = 0x2000 //Declared as an annotation type.
	ACC_ENUM       uint16 = 0x4000 //Declared as an enum type.
)

func (class Class) isInterface() bool {
	return class.AccessFlags&ACC_INTERFACE != uint16(0)
}

func (class Class) isAccessFlagValid() bool {
	if class.isInterface() {
		return class.AccessFlags&ACC_ABSTRACT != uint16(0) &&
			class.AccessFlags&ACC_FINAL == uint16(0) &&
			class.AccessFlags&ACC_SUPER == uint16(0) &&
			class.AccessFlags&ACC_ENUM == uint16(0)
	}
	if class.AccessFlags&ACC_ANNOTATION == uint16(0) {
		return class.AccessFlags&ACC_FINAL != uint16(0) &&
			class.AccessFlags&ACC_ABSTRACT != uint16(0)
	}
	return false
}
