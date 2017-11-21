package classfile

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/ianynchen/gvm/jvm/util"
	"github.com/op/go-logging"
)

/*
 Class file structure
*/
type Class struct {
	Magic            uint32
	MinorVersion     uint16
	MajorVersion     uint16
	ConstantPoolSize uint16
	ConstantPool     [](*ConstantPoolInfo)
	AccessFlags      uint16
	ThisClass        uint16
	SuperClass       uint16
	InterfacesCount  uint16
	Interfaces       []Interface
	FieldsCount      uint16
	Fields           [](*Field)
	MethodsCount     uint16
	Methods          [](*Method)
	AttributesCount  uint16
	Attributes       [](*Attribute)
}

var logger = logging.MustGetLogger("logger")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

type constantPoolHandler = func([]byte) (interface{}, int, error)

var constantPoolMapper = map[uint8]constantPoolHandler{
	CONSTANTClass:              parseConstantClassInfo,
	CONSTANTFieldref:           parseConstantFieldrefInfo,
	CONSTANTMethodref:          parseConstantMethodrefInfo,
	CONSTANTInterfaceMethodref: parseConstantInterfaceMethodrefInfo,
	CONSTANTString:             parseConstantStringInfo,
	CONSTANTInteger:            parseConstantIntegerInfo,
	CONSTANTFloat:              parseConstantFloatInfo,
	CONSTANTLong:               parseConstantLongInfo,
	CONSTANTDouble:             parseConstantDoubleInfo,
	CONSTANTNameAndType:        parseConstantNameAndTypeInfo,
	CONSTANTUtf8:               parseConstantUtf8Info,
	CONSTANTMethodHandle:       parseConstantMethodHandleInfo,
	CONSTANTMethodType:         parseConstantMethodTypeInfo,
	CONSTANTInvokeDynamic:      parseConstantInvokeDynamicInfo,
}

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.DEBUG, "")
	logging.SetBackend(backendLeveled, backendFormatter)
}

/*
NewClass creates a new Class object
*/
func NewClass() *Class {
	return new(Class)
}

func (class Class) Print(writers ...io.Writer) {
	writer := io.MultiWriter(writers...)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "\t")
	encoder.Encode(class)
}

/*
Parse parses Class info from []byte content
*/
func (class *Class) Parse(content []byte, offset int) error {

	// parse the magic number, should always be 0xCAFEBABE
	pos := offset
	magic, err := util.ParseUint32(content[pos:])

	if err != nil {
		return err
	}
	class.Magic = magic
	pos += 4

	// verify the magic number is correct
	logger.Debugf("Magic is %x", class.Magic)
	if class.Magic != 0xCAFEBABE {
		return errors.New("Unexpected magic number")
	}

	// parse minor version
	class.MinorVersion, err = util.ParseUint16(content[pos:])
	pos += 2
	if err != nil {
		return err
	}

	// parse major version
	class.MajorVersion, err = util.ParseUint16(content[pos:])
	pos += 2
	if err != nil {
		return err
	}

	// parse constant pool items
	pos, err = class.parseConstantPool(content, pos)

	class.AccessFlags, err = util.ParseUint16(content[pos:])
	pos += 2
	if err != nil {
		return err
	}

	class.ThisClass, err = util.ParseUint16(content[pos:])
	pos += 2
	if err != nil {
		return err
	}

	class.SuperClass, err = util.ParseUint16(content[pos:])
	pos += 2
	if err != nil {
		return err
	}

	pos, err = class.parseInterfaces(content, pos)
	if err != nil {
		return err
	}
	return err
}

func (class *Class) parseConstantPool(content []byte, offset int) (int, error) {

	// reading constant pool size
	pos := offset
	size, err := util.ParseUint16(content[pos:])
	pos += 2
	if err != nil {
		return pos, err
	}
	class.ConstantPoolSize = size
	if size < 1 {
		return pos, errors.New("constant pool size error")
	}
	logger.Debugf("creating constant pool with size %d", class.ConstantPoolSize)
	class.ConstantPool = make([](*ConstantPoolInfo), class.ConstantPoolSize)

	i := uint16(1)
	for i <= class.ConstantPoolSize-1 {
		tag, err := util.ParseUint8(content[pos:])
		if err != nil {
			return pos + 1, err
		}
		pos++

		class.ConstantPool[i], err = newConstantPool(tag)

		increment := uint16(1)
		if tag == CONSTANTDouble || tag == CONSTANTLong {
			increment = uint16(2)
		}
		if _, ok := constantPoolMapper[tag]; !ok {
			return pos, errors.New("Unable to find constant pool info handler")
		}
		if info, off, err := constantPoolMapper[tag](content[pos:]); err == nil {
			logger.Infof("parsing %d", tag)
			class.ConstantPool[i].Info = info
			class.ConstantPool[i].Position = int(i)
			i += increment
			pos += off
		} else {
			return pos, err
		}
	}
	return pos, nil
}
