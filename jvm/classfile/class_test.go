package classfile

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/issue9/assert"
)

func TestReadHeader(t *testing.T) {

	f, err := os.Open("HelloWorld.class")
	defer f.Close()
	assert.IsNil(err)

	content, err := ioutil.ReadAll(f)

	assert.IsNil(err)
	class := NewClass()
	err = class.Parse(content, 0)
	assert.IsNil(err)
	assert.Equal(t, 0xCAFEBABE, class.Magic)

	class.Print(os.Stdout)
}
