package uniconv

import "time"

func FirstInterface(arr ...interface{}) interface{} {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}
func FirstStr(arr ...string) string {
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}
func FirstByte(arr ...byte) byte {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstRune(arr ...rune) rune {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstInt(arr ...int) int {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstFloat32(arr ...float32) float32 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstInt8(arr ...int8) int8 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstInt16(arr ...int16) int16 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstInt32(arr ...int32) int32 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstInt64(arr ...int64) int64 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstUint8(arr ...uint8) uint8 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstUint(arr ...uint) uint {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstUint16(arr ...uint16) uint16 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstUint32(arr ...uint32) uint32 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstUint64(arr ...uint64) uint64 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstFloat64(arr ...float64) float64 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstDuration(arr ...time.Duration) time.Duration {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
func FirstTime(arr ...time.Time) time.Time {
	if len(arr) > 0 {
		return arr[0]
	}
	return time.Unix(0, 0)
}
func FirstBytes(arr ...[]byte) []byte {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}
func FirstRunes(arr ...[]rune) []rune {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}
func FirstStrings(arr ...[]string) []string {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}
func FirstInts(arr ...[]int) []int {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}
func FirstFloats(arr ...[]float32) []float32 {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}
func FirstChanOfStruct(arr ...chan struct{}) chan struct{} {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}
func FirstChanOfInt(arr ...chan int) chan int {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}
func FirstChanOfStr(arr ...chan string) chan string {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}
func FirstChanOfError(arr ...chan error) chan error {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}
