package algorithm_test

import (
	"github.com/newkedison/go-library/algorithm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackUInt32(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(algorithm.PackUInt32(0), []byte{0x00})
	assert.Equal(algorithm.PackUInt32(1), []byte{0x01})
	assert.Equal(algorithm.PackUInt32(0x7F), []byte{0x7F})
	assert.Equal(algorithm.PackUInt32(0x80), []byte{0x80, 0x80})
	assert.Equal(algorithm.PackUInt32(0xFF), []byte{0x80, 0xFF})
	assert.Equal(algorithm.PackUInt32(0x100), []byte{0xC0, 0x01, 0x00})
	assert.Equal(algorithm.PackUInt32(0xFFF), []byte{0xC0, 0x0F, 0xFF})
	assert.Equal(algorithm.PackUInt32(0xFFFF), []byte{0xC0, 0xFF, 0xFF})
	assert.Equal(algorithm.PackUInt32(0x10000), []byte{0xE0, 0x01, 0x00, 0x00})
	assert.Equal(algorithm.PackUInt32(0xFFFFF), []byte{0xE0, 0x0F, 0xFF, 0xFF})
	assert.Equal(algorithm.PackUInt32(0x1FFFFF), []byte{0xE0, 0x1F, 0xFF, 0xFF})
	assert.Equal(algorithm.PackUInt32(0x200000), []byte{0xE0, 0x20, 0x00, 0x00})
	assert.Equal(algorithm.PackUInt32(0x2FFFFF), []byte{0xE0, 0x2F, 0xFF, 0xFF})
	assert.Equal(algorithm.PackUInt32(0xFFFFFF), []byte{0xE0, 0xFF, 0xFF, 0xFF})
	assert.Equal(algorithm.PackUInt32(0x1000000), []byte{0xF0, 0x01, 0x00, 0x00, 0x00})
	assert.Equal(algorithm.PackUInt32(0xFFFFFFF), []byte{0xF0, 0x0F, 0xFF, 0xFF, 0xFF})
	assert.Equal(algorithm.PackUInt32(0x10000000), []byte{0xF0, 0x10, 0x00, 0x00, 0x00})
	assert.Equal(algorithm.PackUInt32(0x1FFFFFFF), []byte{0xF0, 0x1F, 0xFF, 0xFF, 0xFF})
	assert.Equal(algorithm.PackUInt32(0xFFFFFFFE), []byte{0xF0, 0xFF, 0xFF, 0xFF, 0xFE})
	assert.Equal(algorithm.PackUInt32(0xFFFFFFFF), []byte{0xF0, 0xFF, 0xFF, 0xFF, 0xFF})
}

func getUnPackUInt32Value(data []byte) uint32 {
	v, err := algorithm.UnPackUInt32(data)
	if err != nil {
		panic(err.Error())
	}
	return v
}

func getUnPackUInt32Error(data []byte) error {
	_, err := algorithm.UnPackUInt32(data)
	return err
}

func TestUnPackUInt32(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(getUnPackUInt32Value([]byte{0x00}), uint32(0))
	assert.Equal(getUnPackUInt32Value([]byte{0x7F}), uint32(0x7F))
	assert.Equal(getUnPackUInt32Value([]byte{0x80, 0x80}), uint32(0x80))
	assert.Equal(getUnPackUInt32Value([]byte{0x80, 0xFF}), uint32(0xFF))
	assert.Equal(getUnPackUInt32Value([]byte{0xC0, 0x01, 0x00}), uint32(0x100))
	assert.Equal(getUnPackUInt32Value([]byte{0xC0, 0xFF, 0xFF}), uint32(0xFFFF))
	assert.Equal(getUnPackUInt32Value([]byte{0xE0, 0x01, 0x00, 0x00}), uint32(0x10000))
	assert.Equal(getUnPackUInt32Value([]byte{0xE0, 0xFF, 0xFF, 0xFF}), uint32(0xFFFFFF))
	assert.Equal(getUnPackUInt32Value([]byte{0xF0, 0x01, 0x00, 0x00, 0x00}), uint32(0x1000000))
	assert.Equal(getUnPackUInt32Value([]byte{0xF0, 0xFF, 0xFF, 0xFF, 0xFF}), uint32(0xFFFFFFFF))

	assert.Error(getUnPackUInt32Error([]byte{}))
	assert.Error(getUnPackUInt32Error([]byte{0x80}))
	assert.Error(getUnPackUInt32Error([]byte{0xC0}))
	assert.Error(getUnPackUInt32Error([]byte{0xC0, 0xFF}))
	assert.Error(getUnPackUInt32Error([]byte{0xE0}))
	assert.Error(getUnPackUInt32Error([]byte{0xE0, 0xFF}))
	assert.Error(getUnPackUInt32Error([]byte{0xE0, 0xFF, 0xFF}))
	assert.Error(getUnPackUInt32Error([]byte{0xF0}))
	assert.Error(getUnPackUInt32Error([]byte{0xF0, 0xFF}))
	assert.Error(getUnPackUInt32Error([]byte{0xF0, 0xFF, 0xFF}))
	assert.Error(getUnPackUInt32Error([]byte{0xF0, 0xFF, 0xFF, 0xFF}))

	assert.Error(getUnPackUInt32Error([]byte{0x80, 0x00}))
	assert.Error(getUnPackUInt32Error([]byte{0x80, 0x7F}))
	assert.Nil(getUnPackUInt32Error([]byte{0x80, 0x80}))
	assert.Error(getUnPackUInt32Error([]byte{0xC0, 0x00, 0x00}))
	assert.Error(getUnPackUInt32Error([]byte{0xC0, 0x00, 0xFF}))
	assert.Nil(getUnPackUInt32Error([]byte{0xC0, 0x01, 0x00}))
	assert.Error(getUnPackUInt32Error([]byte{0xE0, 0x00, 0x00, 0x00}))
	assert.Error(getUnPackUInt32Error([]byte{0xE0, 0x00, 0xFF, 0xFF}))
	assert.Nil(getUnPackUInt32Error([]byte{0xE0, 0x01, 0x00, 0x00}))
	assert.Error(getUnPackUInt32Error([]byte{0xF0, 0x00, 0x00, 0x00, 0x00}))
	assert.Error(getUnPackUInt32Error([]byte{0xF0, 0x00, 0xFF, 0xFF, 0xFF}))
	assert.Nil(getUnPackUInt32Error([]byte{0xF0, 0x01, 0x00, 0x00, 0x00}))
}

func TestGetPackedUInt32Length(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(algorithm.GetPackedUInt32Length(0), uint32(1))
	assert.Equal(algorithm.GetPackedUInt32Length(0x7F), uint32(1))
	assert.Equal(algorithm.GetPackedUInt32Length(0x80), uint32(2))
	assert.Equal(algorithm.GetPackedUInt32Length(0xFF), uint32(2))
	assert.Equal(algorithm.GetPackedUInt32Length(0x100), uint32(3))
	assert.Equal(algorithm.GetPackedUInt32Length(0xFFFF), uint32(3))
	assert.Equal(algorithm.GetPackedUInt32Length(0x10000), uint32(4))
	assert.Equal(algorithm.GetPackedUInt32Length(0xFFFFFF), uint32(4))
	assert.Equal(algorithm.GetPackedUInt32Length(0x1000000), uint32(5))
	assert.Equal(algorithm.GetPackedUInt32Length(0xFFFFFFFF), uint32(5))
}
