package libs

import (
	"fmt"
	"time"
)

func GetNextNBits(n int, startBitIndex int, bytes []uint8) uint8 {
	arraySize := len(bytes)
	if startBitIndex < 0 || startBitIndex > (arraySize*8)-1 {
		panic("startBitIndex is too high")
	}

	if n > 0 && n <= 8 {
		res := uint8(0)
		for i := 0; i < n; i++ {
			arrayIndex := (startBitIndex + i) / 8
			bitIndex := (startBitIndex + i) % 8

			bits := bytes[arrayIndex]

			// extract bit at bitIndex location
			mask := uint8(1 << (7 - bitIndex))
			masked := bits & mask
			bit := masked >> (7 - bitIndex)
			res = (res << 1) | bit
		}

		return res
	}

	panic("can only retrieve max 8 bits at the time")
}

func IsBitSet(bitN int, val uint8) bool {
	return (val>>bitN)&1 == 1
}

type IntPoint struct {
	X, Y int
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func TimeTrack(what string) func() {
	start := time.Now()

	return func() {
		elapsed := time.Since(start)
		fmt.Printf("%s took %v\n", what, elapsed)
	}
}
