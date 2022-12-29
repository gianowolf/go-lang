package main

import (
	"encoding/binary"
	"math/rand"
)

const (
	size int64 = 10
)

var (
	boxes []int64
)

func seekRand() *rand.Rand {
	var b [8]byte
	r := rand.New(rand.NewSource(int64(binary.LittleEndian.Uint64(b[:]))))
	return r
}

func main() {
	r := seekRand()

	// init boxes
	boxes = make(int64[], 100)
	for i := 0; i < 100; i++ {

	}
}
