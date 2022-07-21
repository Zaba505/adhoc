package main

import (
	"math/rand"
	"testing"
)

func TestRandInt(t *testing.T) {
	t.Run("should return a different number given a different seed", func(subT *testing.T) {
		low, high := 0, 10
		r1 := rand.New(rand.NewSource(1))
		r2 := rand.New(rand.NewSource(2))

		n1 := randInt(r1, low, high)
		n2 := randInt(r2, low, high)

		if n1 == n2 {
			subT.Fail()
		}
	})
}
