package uniconv

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	. "github.com/rshmelev/go-uniconv/first"
)

type SoftConverter struct {
	Value interface{}
}

var Cast ConverterFunc = func(a interface{}) UniversalConverter {
	x := &SoftConverter{a}
	return x
}

//------------------------

func (c *SoftConverter) Interface() interface{} {
	return c.Value
}
func (c *SoftConverter) Str(fallback ...string) string {
	s, ok := c.Value.(string)
	if ok {
		return s
	}
	return fmt.Sprint(c.Value)
}
func (c *SoftConverter) Byte(fallback ...byte) byte {
	return byte(IntValueWithFallback(c.Value, int64(FirstByte(fallback))))
}
func (c *SoftConverter) Rune(fallback ...rune) rune {
	return rune(IntValueWithFallback(c.Value, int64(FirstRune(fallback))))
}
func (c *SoftConverter) Int(fallback ...int) int {
	return int(IntValueWithFallback(c.Value, int64(FirstInt(fallback))))
}
func (c *SoftConverter) Int8(fallback ...int8) int8 {
	return int8(IntValueWithFallback(c.Value, int64(FirstInt8(fallback))))
}
func (c *SoftConverter) Int16(fallback ...int16) int16 {
	return int16(IntValueWithFallback(c.Value, int64(FirstInt16(fallback))))
}
func (c *SoftConverter) Int32(fallback ...int32) int32 {
	return int32(IntValueWithFallback(c.Value, int64(FirstInt32(fallback))))
}
func (c *SoftConverter) Int64(fallback ...int64) int64 {
	return int64(IntValueWithFallback(c.Value, int64(FirstInt64(fallback))))
}
func (c *SoftConverter) Uint8(fallback ...uint8) uint8 {
	return uint8(IntValueWithFallback(c.Value, int64(FirstUint8(fallback))))
}
func (c *SoftConverter) Uint16(fallback ...uint16) uint16 {
	return uint16(IntValueWithFallback(c.Value, int64(FirstUint16(fallback))))
}
func (c *SoftConverter) Uint32(fallback ...uint32) uint32 {
	return uint32(IntValueWithFallback(c.Value, int64(FirstUint32(fallback))))
}
func (c *SoftConverter) Uint64(fallback ...uint64) uint64 {
	return uint64(IntValueWithFallback(c.Value, int64(FirstUint64(fallback))))
}

func (c *SoftConverter) Float64(fallback ...float64) float64 {
	return float64(FloatValueWithFallback(c.Value, float64(FirstFloat64(fallback))))
}
func (c *SoftConverter) Float32(fallback ...float32) float32 {
	return float32(FloatValueWithFallback(c.Value, float64(FirstFloat32(fallback))))
}

func (c *SoftConverter) Duration(fallback ...time.Duration) time.Duration {
	return time.Duration(IntValueWithFallback(c.Value, int64(FirstDuration(fallback))))
}
func (c *SoftConverter) Time(fallback ...time.Time) time.Time {
	if v, ok := TimeValue(c.Value); ok {
		return v
	}
	return FirstTime(fallback)
}

func (c *SoftConverter) Bytes(fallback ...[]byte) []byte {
	a := c.Value
	if a != nil {
		switch n := a.(type) {
		case string:
			return []byte(n)
		case []rune:
			return []byte(string(n))
		case []byte:
			return n
		}
	}
	return FirstBytes(fallback)
}
func (c *SoftConverter) Runes(fallback ...[]rune) []rune {
	a := c.Value
	if a != nil {
		switch n := a.(type) {
		case string:
			return []rune(n)
		case []rune:
			return n
		case []byte:
			return []rune(string(n))
		}
	}
	return FirstRunes(fallback)

}
func (c *SoftConverter) Strings(fallback ...[]string) []string {
	a := c.Value
	if a != nil {
		switch n := a.(type) {
		case []string:
			return n
		case string:
			return []string{n}
		case []byte:
			return []string{string(n)}
		case []int:
			res := make([]string, len(n))
			for i := 0; i < len(n); i++ {
				res[i] = strconv.Itoa(n[i])
			}
			return res
		case []rune:
			return []string{string(n)}
		}
	}
	return FirstStrings(fallback)
}
func (c *SoftConverter) Ints(fallback ...[]int) []int {
	a := c.Value
	if a != nil {
		switch n := a.(type) {
		case []string:
			res := make([]int, len(n))
			for i := 0; i < len(n); i++ {
				res[i], _ = strconv.Atoi(n[i])
			}
			return res

		case []float32:
			res := make([]int, len(n))
			for i := 0; i < len(n); i++ {
				res[i] = int(n[i])
			}
			return res

		case []float64:
			res := make([]int, len(n))
			for i := 0; i < len(n); i++ {
				res[i] = int(n[i])
			}
			return res

		case []int:
			return n
		}
	}
	return FirstInts(fallback)
}
func (c *SoftConverter) Floats32(fallback ...[]float32) []float32 {
	a := c.Value
	if a != nil {
		switch n := a.(type) {
		case []float32:
			return n
		// todo
		case []string:
			res := make([]float32, len(n))
			for i := 0; i < len(n); i++ {
				t, _ := strconv.ParseFloat(n[i], 32)
				res[i] = float32(t)
			}
			return res
		case []int:
			res := make([]float32, len(n))
			for i := 0; i < len(n); i++ {
				res[i] = float32(n[i])
			}
			return res
		case []float64:
			res := make([]float32, len(n))
			for i := 0; i < len(n); i++ {
				res[i] = float32(n[i])
			}
			return res

		}
	}
	return FirstFloats32(fallback)
}
func (c *SoftConverter) Floats64(fallback ...[]float64) []float64 {
	a := c.Value
	if a != nil {
		switch n := a.(type) {
		case []float64:
			return n
		// todo
		case []string:
			res := make([]float64, len(n))
			for i := 0; i < len(n); i++ {
				res[i], _ = strconv.ParseFloat(n[i], 64)
			}
			return res
		case []int:
			res := make([]float64, len(n))
			for i := 0; i < len(n); i++ {
				res[i] = float64(n[i])
			}
			return res
		case []float32:
			res := make([]float64, len(n))
			for i := 0; i < len(n); i++ {
				res[i] = float64(n[i])
			}
			return res
		}
	}
	return FirstFloats64(fallback)
}

//---------------------------------------------------

func IntValueWithFallback(a interface{}, fallback int64) int64 {
	x, ok := IntValue(a)
	if !ok {
		return fallback
	}
	return x
}
func FloatValueWithFallback(a interface{}, fallback float64) float64 {
	x, ok := FloatValue(a)
	if !ok {
		return fallback
	}
	return x
}

func FloatValue(a interface{}) (float64, bool) {
	if a == nil {
		return 0, false
	}
	switch n := a.(type) {
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case int64:
		return float64(n), true
	case time.Duration:
		return float64(n.Seconds()), true
	case int:
		return float64(n), true
	case int8:
		return float64(n), true
	case int16:
		return float64(n), true
	case int32:
		return float64(n), true
	case uint:
		return float64(n), true
	case uint8:
		return float64(n), true
	case uint16:
		return float64(n), true
	case uint32:
		return float64(n), true
	case uint64:
		return float64(n), true
	case float64:
		return float64(n), true
	case float32:
		return float64(n), true
	case string:
		x, err := strconv.ParseFloat(n, 64)
		return x, err == nil
	default:
		return 0, false
	}
}

func IntMsToTime(a int64) time.Time {
	return time.Unix(0, a*int64(time.Millisecond))
}

func TimeValue(a interface{}) (time.Time, bool) {
	if a == nil {
		return time.Unix(0, 0), false
	}
	switch n := a.(type) {
	case int64:
		return IntMsToTime(int64(n)), true
	case time.Duration:
		return time.Now().Add(n), true
	case int:
		return IntMsToTime(int64(n) * 1000), true
	case int32:
		return IntMsToTime(int64(n) * 1000), true
	case uint:
		return IntMsToTime(int64(n) * 1000), true
	case uint32:
		return IntMsToTime(int64(n) * 1000), true
	case uint64:
		return IntMsToTime(int64(n)), true
	case float64:
		return IntMsToTime(int64(n * 1000)), true
	case float32:
		return IntMsToTime(int64(n)), true
	case string:
		x, err := ParseTime(n)
		return x, err == nil
	default:
		return time.Unix(0, 0), false
	}
}

func IntValue(a interface{}) (int64, bool) {
	if a == nil {
		return 0, false
	}

	switch n := a.(type) {
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	case int64:
		return int64(n), true
	case time.Duration:
		return int64(n), true
	case int:
		return int64(n), true
	case int8:
		return int64(n), true
	case int16:
		return int64(n), true
	case int32:
		return int64(n), true
	case uint:
		return int64(n), true
	case uint8:
		return int64(n), true
	case uint16:
		return int64(n), true
	case uint32:
		return int64(n), true
	case uint64:
		return int64(n), true
	case float64:
		return int64(n), true
	case float32:
		return int64(n), true
	case string:
		x, err := strconv.ParseInt(n, 10, 64)
		return int64(x), err == nil
	default:
		return 0, false
	}
}

var TimeFormats = []string{"1 2 2006", "1 2 2006 15 4 5", "2006 1 2 15 4 5", "2006 1 2", "15 4 5 Jan 2, 2006 MST"}

var replacer, _ = regexp.Compile("[/,. \\-:]+")
var replacerError = errors.New("Can't parse string as time")

func ParseTime(str string) (t time.Time, err error) {
	str = replacer.ReplaceAllString(str, " ")
	for _, format := range TimeFormats {
		t, err = time.Parse(format, str)
		if err == nil {
			return
		}
	}
	err = replacerError
	return
}
