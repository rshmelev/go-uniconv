package uniconv

import "time"

type UniversalConverter interface {
	Interface() interface{}

	Str(fallback ...string) string

	Byte(fallback ...byte) byte
	Rune(fallback ...rune) rune
	Int(fallback ...int) int
	Float32(fallback ...float32) float32

	Int8(fallback ...int8) int8
	Int16(fallback ...int16) int16
	Int32(fallback ...int32) int32
	Int64(fallback ...int64) int64

	Uint8(fallback ...uint8) uint8
	Uint16(fallback ...uint16) uint16
	Uint32(fallback ...uint32) uint32
	Uint64(fallback ...uint64) uint64

	Float64(fallback ...float64) float64

	Duration(fallback ...time.Duration) time.Duration
	Time(fallback ...time.Time) time.Time

	Bytes(fallback ...[]byte) []byte
	Runes(fallback ...[]rune) []rune
	Strings(fallback ...[]string) []string
	Ints(fallback ...[]int) []int
	Floats32(fallback ...[]float32) []float32
	Floats64(fallback ...[]float64) []float64
}

type ConverterFunc func(interface{}) UniversalConverter
