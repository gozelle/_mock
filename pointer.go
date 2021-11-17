package _mock

func ToStringPointer(v string) *string {
	r := new(string)
	*r = v
	return r
}

func ToIntPointer(v int) *int {
	r := new(int)
	*r = v
	return r
}

func ToInt8Pointer(v int8) *int8 {
	r := new(int8)
	*r = v
	return r
}

func ToInt16Pointer(v int16) *int16 {
	r := new(int16)
	*r = v
	return r
}

func ToInt32Pointer(v int32) *int32 {
	r := new(int32)
	*r = v
	return r
}

func ToInt64Pointer(v int64) *int64 {
	r := new(int64)
	*r = v
	return r
}

func ToUintPointer(v uint) *uint {
	r := new(uint)
	*r = v
	return r
}

func ToUint8Pointer(v uint8) *uint8 {
	r := new(uint8)
	*r = v
	return r
}

func ToUint16Pointer(v uint16) *uint16 {
	r := new(uint16)
	*r = v
	return r
}

func ToUint32Pointer(v uint32) *uint32 {
	r := new(uint32)
	*r = v
	return r
}

func ToUint64Pointer(v uint64) *uint64 {
	r := new(uint64)
	*r = v
	return r
}

func ToFloat32Pointer(v float32) *float32 {
	r := new(float32)
	*r = v
	return r
}

func ToFloat64Pointer(v float64) *float64 {
	r := new(float64)
	*r = v
	return r
}
