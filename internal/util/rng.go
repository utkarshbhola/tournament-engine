package util

import (
	"math/rand"
	"time"
)

func NewRng() *rand.Rand {
	seed := time.Now().UnixNano()
	return rand.New(rand.NewSource(seed))
}
