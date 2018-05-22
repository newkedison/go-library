package algorithm

import (
	"errors"
)

// PackUInt32 convert an uint32 value to several bytes
func PackUInt32(v uint32) []byte {
	switch {
	case v <= 0x7F:
		return []byte{byte(v & 0xFF)}
	case v <= 0xFF:
		return []byte{0x80, byte(v & 0xFF)}
	case v <= 0xFFFF:
		return []byte{0xC0, byte((v >> 8) & 0xFF), byte(v & 0xFF)}
	case v <= 0xFFFFFF:
		return []byte{
			0xE0, byte((v >> 16) & 0xFF), byte((v >> 8) & 0xFF), byte(v & 0xFF)}
	default:
		return []byte{0xF0, byte((v >> 24) & 0xFF),
			byte((v >> 16) & 0xFF), byte((v >> 8) & 0xFF), byte(v & 0xFF)}
	}
}

// GetPackedUInt32Length return the byte count of v when it was packed
func GetPackedUInt32Length(v uint32) uint32 {
	switch {
	case v <= 0x7F:
		return 1
	case v <= 0xFF:
		return 2
	case v <= 0xFFFF:
		return 3
	case v <= 0xFFFFFF:
		return 4
	default:
		return 5
	}
}

// UnPackUInt32 convert the first several bytes to an uint32 if it match the format
func UnPackUInt32(data []byte) (uint32, error) {
	var v uint32
	l := len(data)
	if l == 0 {
		return v, errors.New("PackedUInt.UnPackUInt32 fail: empty data")
	}
	if data[0]&0x80 == 0 {
		v = uint32(data[0] & 0x7F)
	} else if l > 1 && data[0] == 0x80 && data[1] > 0x7F {
		v = uint32(data[1])
	} else if l > 2 && data[0] == 0xC0 && data[1] > 0 {
		v = uint32((uint32(data[1]) << 8) | uint32(data[2]))
	} else if l > 3 && data[0] == 0xE0 && data[1] > 0 {
		v = uint32((uint32(data[1]) << 16) |
			(uint32(data[2]) << 8) | uint32(data[3]))
	} else if l > 4 && data[0] == 0xF0 && data[1] > 0 {
		v = uint32((uint32(data[1]) << 24) | (uint32(data[2]) << 16) |
			(uint32(data[3]) << 8) | uint32(data[4]))
	} else {
		return v, errors.New("PackedUInt.UnPackUInt32 fail: unknown format")
	}
	return v, nil
}
