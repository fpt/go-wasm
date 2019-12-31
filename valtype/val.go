package valtype

import (
)

type WAType int

const (
	I32 WAType = iota
	I64 WAType = iota
	F32 WAType = iota
	F64 WAType = iota
)

func ValType(b byte) WAType {
	switch b {
	case 0x7e:
		return I64
	case 0x7d:
		return F32
	case 0x7c:
		return F64
	}
	// 0x7f
	return I32
}
