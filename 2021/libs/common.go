package libs

import (
	"fmt"
	"time"
)

func IsBitSet(bitN int, val uint64) bool {
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
