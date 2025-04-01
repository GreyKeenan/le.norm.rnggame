package toddler

import (
	"math/rand"
)

func Control(r *rand.Rand, games int) float64 {
	var total int64
	var n, h int
	for i := 0; i < games; i++ {
		h = 0
		for {
			n = rval(r)
			if n == maximum { break }
			if n > h { h = n }
			if r.Int() & 1 != 0 {
				total += int64(h)
				break
			}
		}
	}
	return float64(total) / float64(games)
}
