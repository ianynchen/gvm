package classfile

import (
	"errors"
	"os"

	"github.com/ianynchen/gvm/jvm/util"
	"github.com/op/go-logging"
)

/*
 Class file structure
*/
type Class struct {
	Magic        uint32
	MinorVersion uint16
	MajorVersion uint16
}

var logger = logging.MustGetLogger("logger")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.DEBUG, "")
	logging.SetBackend(backendLeveled, backendFormatter)
}

func NewClass() *Class {
	return new(Class)
}

/*
 Parses Class info from []byte content
*/
func (class *Class) Parse(content []byte, offset int) error {

	pos := offset
	magic, err := util.ParseUint32(content[pos:])

	if err != nil {
		return err
	}
	class.Magic = magic
	pos += 4

	logger.Debugf("Magic is %x", class.Magic)
	if class.Magic != 0xCAFEBABE {
		return errors.New("Unexpected magic number")
	}

	class.MinorVersion, err = util.ParseUint16(content[pos:])
	pos += 2

	if err != nil {
		return err
	}

	class.MajorVersion, err = util.ParseUint16(content[pos:])
	pos += 2

	if err != nil {
		return err
	}
	return nil
}
