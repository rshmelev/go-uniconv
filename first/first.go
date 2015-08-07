package uniconv

import "time"

func FirstInterface(arr []interface{}, alt ...interface{}) interface{} {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
func FirstStr(arr []string, alt ...string) string {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return ""
}
func FirstByte(arr []byte, alt ...byte) byte {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstRune(arr []rune, alt ...rune) rune {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstInt(arr []int, alt ...int) int {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstFloat32(arr []float32, alt ...float32) float32 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstInt8(arr []int8, alt ...int8) int8 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstInt16(arr []int16, alt ...int16) int16 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstInt32(arr []int32, alt ...int32) int32 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstInt64(arr []int64, alt ...int64) int64 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstUint8(arr []uint8, alt ...uint8) uint8 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstUint(arr []uint, alt ...uint) uint {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstUint16(arr []uint16, alt ...uint16) uint16 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstUint32(arr []uint32, alt ...uint32) uint32 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstUint64(arr []uint64, alt ...uint64) uint64 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstFloat64(arr []float64, alt ...float64) float64 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstDuration(arr []time.Duration, alt ...time.Duration) time.Duration {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return 0
}
func FirstTime(arr []time.Time, alt ...time.Time) time.Time {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return time.Unix(0, 0)
}
func FirstBytes(arr [][]byte, alt ...[]byte) []byte {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
func FirstRunes(arr [][]rune, alt ...[]rune) []rune {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
func FirstStrings(arr [][]string, alt ...[]string) []string {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
func FirstInts(arr [][]int, alt ...[]int) []int {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
func FirstFloats32(arr [][]float32, alt ...[]float32) []float32 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
func FirstFloats64(arr [][]float64, alt ...[]float64) []float64 {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
func FirstChanOfStruct(arr []chan struct{}, alt ...chan struct{}) chan struct{} {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
func FirstChanOfInt(arr []chan int, alt ...chan int) chan int {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
func FirstChanOfStr(arr []chan string, alt ...chan string) chan string {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
func FirstChanOfError(arr []chan error, alt ...chan error) chan error {
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	if len(alt) > 0 {
		return alt[0]
	}
	return nil
}
