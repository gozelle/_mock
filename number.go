package _mock

import (
	"crypto/rand"
	"encoding/binary"
	"github.com/gozelle/decimal"
	"math"
	"math/big"
	mathRand "math/rand"
	"strconv"
)

func Int() int {
	if strconv.IntSize == 32 {
		return int(Int32())
	}
	return int(Int64())
}

func IntMin() int {
	if strconv.IntSize == 32 {
		return math.MinInt32
	}
	return math.MinInt64
}

func IntMax() int {
	if strconv.IntSize == 32 {
		return math.MaxInt32
	}
	return math.MaxInt64
}

func IntRange(min, max int) int {
	return int(Int64Range(int64(min), int64(max)))
}

func Int8() int8 {
	return int8(Int64())
}

func Int8Min() int8 {
	return math.MinInt8
}

func Int8Max() int8 {
	return math.MaxInt8
}

func Int8Range(min, max int8) int8 {
	return int8(Int64Range(int64(min), int64(max)))
}

func Int16() int16 {
	return int16(Int64())
}

func Int16Min() int16 {
	return math.MinInt16
}

func Int16Max() int16 {
	return math.MaxInt16
}

func Int16Range(min, max int16) int16 {
	return int16(Int64Range(int64(min), int64(max)))
}

func Int32() int32 {
	return int32(Int64())
}

func Int32Min() int32 {
	return math.MinInt32
}

func Int32Max() int32 {
	return math.MaxInt32
}

func Int32Range(min, max int32) int32 {
	return int32(Int64Range(int64(min), int64(max)))
}

func Int64() int64 {
	return Int64Range(math.MinInt64, math.MaxInt64)
}

func Int64Min() int64 {
	return math.MinInt64
}

func Int64Max() int64 {
	return math.MaxInt64
}

func Int64Range(min, max int64) int64 {
	if min == max {
		return min
	}
	if min > max {
		max, min = min, max
	}
	result, err := rand.Int(rand.Reader, big.NewInt(0).Sub(big.NewInt(max), big.NewInt(min)))
	if err != nil {
		return 0
	}
	r := result.Int64() + min
	
	b, err := rand.Int(rand.Reader, big.NewInt(2))
	if err == nil && b.Int64() == 1 {
		r += 1
	}
	return r
}

func Uint() uint {
	if strconv.IntSize == 32 {
		return uint(Uint32())
	}
	return uint(Uint64())
}

func UintMin() uint {
	return 0
}

func UintMax() uint {
	return 0
}

func UintRange(min, max uint) uint {
	return uint(Uint64Range(uint64(min), uint64(max)))
}

func Uin8() uint8 {
	return uint8(Uint64Range(0, math.MaxUint16))
}

func Uin8Min() uint8 {
	return 0
}

func Uin8Max() uint8 {
	return math.MaxUint8
}

func Uint8Range(min, max uint8) uint8 {
	return uint8(Uint64Range(uint64(min), uint64(max)))
}

func Uint16() uint16 {
	return uint16(Uint64Range(0, math.MaxUint16))
}

func Uin16tMin() uint16 {
	return 0
}

func Uin16tMax() uint16 {
	return math.MaxUint16
}

func Uint16Range(min, max uint16) uint16 {
	return uint16(Uint64Range(uint64(min), uint64(max)))
}

func Uint32() uint32 {
	return uint32(Uint64Range(0, math.MaxUint32))
}

func Uint32Min() uint32 {
	return 0
}

func Uint32Max() uint32 {
	return math.MaxUint32
}

func Uint32Range(min, max uint32) uint32 {
	return uint32(Uint64Range(uint64(min), uint64(max)))
}

func Uint64() uint64 {
	return Uint64Range(0, math.MaxUint64)
}

func Uint64Max() uint64 {
	return math.MaxUint64
}

func Uint64Range(min, max uint64) uint64 {
	if min <= math.MaxInt64 && max <= math.MaxInt64 {
		return uint64(Int64Range(int64(min), int64(max)))
	}
	buf := make([]byte, 8)
	_, _ = rand.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}

func Float32() float32 {
	return mathRand.Float32()
}

func Float64() float64 {
	return mathRand.Float64()
}

func Decimal() decimal.Decimal {
	return decimal.NewFromFloat(Float64())
}


