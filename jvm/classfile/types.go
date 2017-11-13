package classfile

const (
	PrimitiveTypeBoolean = iota // unknown length
	// newarray enables creation of boolean arrays
	// arrays of type boolean are accessed and modified using baload and bastore
	// Oracle JVM uses 8 bits per boolean element

	PrimitiveTypeByte          // 8 bits
	PrimitiveTypeShort         // 16 bits
	PrimitiveTypeInt           // 32 bits
	PrimitiveTypeLong          // 64 bits
	PrimitiveTypeChar          // 16 bits
	PrimitiveTypeFloat         // 32 bits
	PrimitiveTypeDouble        // 64 bits
	PrimitiveTypeReturnAddress // undetermined

	ReferenceTypeClass     // class reference type
	ReferenceTypeArray     // array reference type
	ReferenceTypeInterface // interface reference type
)
