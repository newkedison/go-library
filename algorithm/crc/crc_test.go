package crc_test

import (
	"github.com/newkedison/go-library/algorithm/crc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = makeTestData(256)

func makeTestData(n int) []byte {
	ret := make([]byte, n, n)
	for i := 0; i < n; i++ {
		ret[i] = byte(i)
	}
	return ret
}

func TestCrc8(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(crc.Crc8([]byte{0x10}), byte(0x70))
	assert.Equal(crc.Crc8([]byte{0x10, 0x20, 0x30, 0xFF}), byte(0x2E))
	assert.Equal(crc.Crc8(testData), byte(0x14))
}

func TestCrc8Impl(t *testing.T) {
	assert := assert.New(t)
	table := crc.MakeCrc8Table(0x1D)
	assert.Equal(
		crc.Crc8Impl([]byte{0x10}, table, 0xFD, false, false, 0x00),
		byte(0x33))
	assert.Equal(
		crc.Crc8Impl(testData, table, 0xFD, false, true, 0x00),
		byte(0x03))
}

func TestCrc8Predefined(t *testing.T) {
	assert := assert.New(t)
	if _, err := crc.Crc8Predefined(nil, 100); assert.NotNil(err) {
		assert.EqualError(err, "core/algorithm/crc: invalid index: 100")
	}
	_crc, _ := crc.Crc8Predefined(testData, 0)
	assert.EqualValues(_crc, 0x14)
	_crc, _ = crc.Crc8Predefined(testData, 1)
	assert.EqualValues(_crc, 0x41)
	_crc, _ = crc.Crc8Predefined(testData, 2)
	assert.EqualValues(_crc, 0x3C)
	_crc, _ = crc.Crc8Predefined(testData, 3)
	assert.EqualValues(_crc, 0xCA)
	_crc, _ = crc.Crc8Predefined(testData, 4)
	assert.EqualValues(_crc, 0xC5)
	_crc, _ = crc.Crc8Predefined(testData, 5)
	assert.EqualValues(_crc, 0xC0)
	_crc, _ = crc.Crc8Predefined(testData, 6)
	assert.EqualValues(_crc, 0x41)
	_crc, _ = crc.Crc8Predefined(testData, 7)
	assert.EqualValues(_crc, 0x18)
	_crc, _ = crc.Crc8Predefined(testData, 8)
	assert.EqualValues(_crc, 0x8E)
	_crc, _ = crc.Crc8Predefined(testData, 9)
	assert.EqualValues(_crc, 0x59)
}

func TestCrc8Continue(t *testing.T) {
	_crc := crc.Crc8(testData[:100])
	_crc, _ = crc.Crc8Continue(testData[100:], _crc)
	assert.Equal(t, _crc, byte(0x14))
	_crc = crc.Crc8(testData[:150])
	_crc, _ = crc.Crc8Continue(testData[150:], _crc)
	assert.Equal(t, _crc, byte(0x14))
}

func TestAppendCrc8(t *testing.T) {
	assert := assert.New(t)
	data := append([]byte(nil), testData...)
	data = crc.AppendCrc8(data)
	assert.Equal(len(data), 257)
	assert.Equal(data[256], byte(0x14))
}

func TestVerifyCrc8(t *testing.T) {
	assert := assert.New(t)
	data := append([]byte(nil), testData...)
	data = crc.AppendCrc8(data)
	assert.True(crc.VerifyCrc8(data))
	data[256] ^= 0xFF
	assert.False(crc.VerifyCrc8(data))
}

func TestCrc16(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(crc.Crc16([]byte{0x10}), uint16(0x8CBE))
	assert.Equal(crc.Crc16(testData), uint16(0xDE6C))
}

func TestCrc16Impl(t *testing.T) {
	assert := assert.New(t)
	table := crc.MakeCrc16Table(0x1021)
	assert.Equal(
		crc.Crc16Impl([]byte{0x10}, table, 0x0000, false, false, 0x0000),
		uint16(0x1231))
	assert.Equal(
		crc.Crc16Impl(testData, table, 0xFFFF, false, false, 0x00),
		uint16(0x3FBD))
}

func TestCrc16Predefined(t *testing.T) {
	assert := assert.New(t)
	if _, err := crc.Crc16Predefined(nil, 100); assert.NotNil(err) {
		assert.EqualError(err, "core/algorithm/crc: invalid index: 100")
	}
	_crc, _ := crc.Crc16Predefined(testData, 0)
	assert.EqualValues(_crc, 0x3FBD)
	_crc, _ = crc.Crc16Predefined(testData, 1)
	assert.EqualValues(_crc, 0xBAD3)
	_crc, _ = crc.Crc16Predefined(testData, 2)
	assert.EqualValues(_crc, 0x3C8E)
	_crc, _ = crc.Crc16Predefined(testData, 3)
	assert.EqualValues(_crc, 0x3B7A)
	_crc, _ = crc.Crc16Predefined(testData, 4)
	assert.EqualValues(_crc, 0xF8E6)
	_crc, _ = crc.Crc16Predefined(testData, 5)
	assert.EqualValues(_crc, 0xB5A1)
	_crc, _ = crc.Crc16Predefined(testData, 6)
	assert.EqualValues(_crc, 0x9A1C)
	_crc, _ = crc.Crc16Predefined(testData, 7)
	assert.EqualValues(_crc, 0x9A1D)
	_crc, _ = crc.Crc16Predefined(testData, 8)
	assert.EqualValues(_crc, 0x4472)
	_crc, _ = crc.Crc16Predefined(testData, 9)
	assert.EqualValues(_crc, 0xB50D)
	_crc, _ = crc.Crc16Predefined(testData, 10)
	assert.EqualValues(_crc, 0xC042)
	_crc, _ = crc.Crc16Predefined(testData, 11)
	assert.EqualValues(_crc, 0x452C)
	_crc, _ = crc.Crc16Predefined(testData, 12)
	assert.EqualValues(_crc, 0xCFC3)
	_crc, _ = crc.Crc16Predefined(testData, 13)
	assert.EqualValues(_crc, 0x563B)
	_crc, _ = crc.Crc16Predefined(testData, 14)
	assert.EqualValues(_crc, 0xE0B5)
	_crc, _ = crc.Crc16Predefined(testData, 15)
	assert.EqualValues(_crc, 0x2193)
	_crc, _ = crc.Crc16Predefined(testData, 16)
	assert.EqualValues(_crc, 0xD841)
	_crc, _ = crc.Crc16Predefined(testData, 17)
	assert.EqualValues(_crc, 0xDE6C)
	_crc, _ = crc.Crc16Predefined(testData, 18)
	assert.EqualValues(_crc, 0x303C)
	_crc, _ = crc.Crc16Predefined(testData, 19)
	assert.EqualValues(_crc, 0x7E55)
}

func TestAppendCrc16(t *testing.T) {
	assert := assert.New(t)
	data := append([]byte(nil), testData...)
	data = crc.AppendCrc16(data)
	assert.Equal(len(data), 258)
	assert.Equal(data[256], byte(0x6C))
	assert.Equal(data[257], byte(0xDE))
}

func TestCrc16Continue(t *testing.T) {
	_crc := crc.Crc16(testData[:100])
	_crc, _ = crc.Crc16Continue(testData[100:], _crc)
	assert.Equal(t, _crc, uint16(0xDE6C))
	_crc = crc.Crc16(testData[:150])
	_crc, _ = crc.Crc16Continue(testData[150:], _crc)
	assert.Equal(t, _crc, uint16(0xDE6C))
}

func TestVerifyCrc16(t *testing.T) {
	assert := assert.New(t)
	data := append([]byte(nil), testData...)
	data = crc.AppendCrc16(data)
	assert.True(crc.VerifyCrc16(data))
	data[256] ^= 0xFF
	assert.False(crc.VerifyCrc16(data))
}

func TestCrc32(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(crc.Crc32([]byte{0x10}), uint32(0xCFB5FFE9))
	assert.Equal(crc.Crc32(testData), uint32(0x29058C73))
}

func TestCrc32Impl(t *testing.T) {
	assert := assert.New(t)
	table := crc.MakeCrc32Table(0xA833982B)
	assert.Equal(
		crc.Crc32Impl([]byte{0x10}, table, 0xFFFFFFFF, false, false, 0xFFFFFFFF),
		uint32(0x0E44E111))
	assert.Equal(
		crc.Crc32Impl(testData, table, 0xFFFFFFFF, true, true, 0xFFFFFFFF),
		uint32(0x6165003E))
}

func TestCrc32Predefined(t *testing.T) {
	assert := assert.New(t)
	if _, err := crc.Crc32Predefined(nil, 100); assert.NotNil(err) {
		assert.EqualError(err, "core/algorithm/crc: invalid index: 100")
	}
	_crc, _ := crc.Crc32Predefined(testData, 0)
	assert.EqualValues(_crc, 0x29058C73)
	_crc, _ = crc.Crc32Predefined(testData, 1)
	assert.EqualValues(_crc, 0xB6B5EE95)
	_crc, _ = crc.Crc32Predefined(testData, 2)
	assert.EqualValues(_crc, 0x9C44184B)
	_crc, _ = crc.Crc32Predefined(testData, 3)
	assert.EqualValues(_crc, 0x6165003E)
	_crc, _ = crc.Crc32Predefined(testData, 4)
	assert.EqualValues(_crc, 0x494A116A)
	_crc, _ = crc.Crc32Predefined(testData, 5)
	assert.EqualValues(_crc, 0x53EB78DA)
	_crc, _ = crc.Crc32Predefined(testData, 6)
	assert.EqualValues(_crc, 0xAB16EA85)
	_crc, _ = crc.Crc32Predefined(testData, 7)
	assert.EqualValues(_crc, 0xD6FA738C)
	_crc, _ = crc.Crc32Predefined(testData, 8)
	assert.EqualValues(_crc, 0x3FF2756B)
}

func TestAppendCrc32(t *testing.T) {
	assert := assert.New(t)
	data := append([]byte(nil), testData...)
	data = crc.AppendCrc32(data)
	assert.Equal(len(data), 260)
	assert.Equal(data[256], byte(0x73))
	assert.Equal(data[257], byte(0x8C))
	assert.Equal(data[258], byte(0x05))
	assert.Equal(data[259], byte(0x29))
}

func TestCrc32Continue(t *testing.T) {
	_crc := crc.Crc32(testData[:100])
	_crc, _ = crc.Crc32Continue(testData[100:], _crc)
	assert.Equal(t, _crc, uint32(0x29058C73))
	_crc = crc.Crc32(testData[:150])
	_crc, _ = crc.Crc32Continue(testData[150:], _crc)
	assert.Equal(t, _crc, uint32(0x29058C73))
}

func TestVerifyCrc32(t *testing.T) {
	assert := assert.New(t)
	data := append([]byte(nil), testData...)
	data = crc.AppendCrc32(data)
	assert.True(crc.VerifyCrc32(data))
	data[256] ^= 0xFF
	assert.False(crc.VerifyCrc32(data))
}

func TestChangeCrc8ConfigIndex(t *testing.T) {
	assert := assert.New(t)
	crc.DefaultCrc8ConfigIndex = 100
	assert.Panics(func() { crc.Crc8([]byte{0x00}) })
	crc.DefaultCrc16ConfigIndex = 100
	assert.Panics(func() { crc.Crc16([]byte{0x00}) })
	crc.DefaultCrc32ConfigIndex = 100
	assert.Panics(func() { crc.Crc32([]byte{0x00}) })

	_, err := crc.Crc8Continue([]byte{0x10}, 0)
	assert.Error(err)
	_, err = crc.Crc16Continue([]byte{0x10}, 0)
	assert.Error(err)
	_, err = crc.Crc32Continue([]byte{0x10}, 0)
	assert.Error(err)

	crc.Crc8Configs = append(
		crc.Crc8Configs, crc.Config{0x9B, 0x00, 0x00, true, true})
	crc.DefaultCrc8ConfigIndex = len(crc.Crc8Configs) - 1
	if v, err := crc.Crc8Continue([]byte{0x10},
		byte(crc.Crc8Configs[crc.DefaultCrc8ConfigIndex].InitValue)); assert.Nil(err) {
		assert.EqualValues(v, 0x98)
	}

	crc.Crc16Configs = append(
		crc.Crc16Configs, crc.Config{0x1021, 0x0000, 0x0000, false, false})
	crc.DefaultCrc16ConfigIndex = len(crc.Crc16Configs) - 1
	if v, err := crc.Crc16Continue([]byte{0x10}, uint16(
		crc.Crc16Configs[crc.DefaultCrc16ConfigIndex].InitValue)); assert.Nil(err) {
		assert.EqualValues(v, 0x1231)
	}

	crc.Crc32Configs = append(
		crc.Crc32Configs, crc.Config{0xAF, 0, 0, false, false})
	crc.DefaultCrc32ConfigIndex = len(crc.Crc32Configs) - 1
	if v, err := crc.Crc32Continue([]byte{0x10}, uint32(
		crc.Crc32Configs[crc.DefaultCrc32ConfigIndex].InitValue)); assert.Nil(err) {
		assert.EqualValues(v, 0xAF0)
	}
}
