package util

import (
	"testing"

	"github.com/issue9/assert"
)

func TestParseUint8(t *testing.T) {

	content := []byte{0xCA}
	val, err := ParseUint8(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint8(0xCA))
}

func TestParseUint16(t *testing.T) {

	content := []byte{0xCA, 0xFE}
	val, err := ParseUint16(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint16(0xCAFE))
}

func TestParseUint32(t *testing.T) {

	content := []byte{0xCA, 0xFE, 0xBA, 0xBE}
	val, err := ParseUint32(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint32(0xCAFEBABE))
}

func TestParseUint64(t *testing.T) {

	content := []byte{0xCA, 0xFE, 0xBA, 0xBE, 0x00, 0x00, 0x00, 0x00}
	val, err := ParseUint64(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint64(0xCAFEBABE00000000))
}
