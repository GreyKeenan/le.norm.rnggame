package toddler

import (
	"math/rand"
)

const minimum int = 1
const maximum int = 21
const Possibilities int = maximum - minimum

func rval(r *rand.Rand) int {
	return minimum + r.Intn(Possibilities + 1)
}

func Train(counts *[Possibilities]int, totals *[Possibilities]int, r *rand.Rand, rounds int) {

	var state int
	var curr int

	for i := 0; i < rounds; i++ {
		state = (i % Possibilities) + minimum
		curr = rval(r)

		counts[state - minimum] += 1

		if curr == maximum {
			totals[state - minimum] -= state
		} else if curr > state {
			totals[state - minimum] += curr - state
		}

	}
}

func Solidify(counts *[Possibilities]int, totals *[Possibilities]int, averages *[Possibilities]float64) {
	for i := 0; i < Possibilities; i++ {
		averages[i] = float64(totals[i]) / float64(counts[i])
	}
}

func Play(averages *[Possibilities]float64, r *rand.Rand) int {
	var n int
	var h int = 0
	for {
		n = rval(r)
		if n <= h {
			continue
		}
		if n == maximum {
			return 0
		}
		h = n
		if averages[h - minimum] <= 0 {
			return h
		}
	}
}

func Perform(averages *[Possibilities]float64, r *rand.Rand, games int) float64 {
	var total int64
	for i := 0; i < games; i++ {
		total += int64(Play(averages, r))
	}
	return float64(total) / float64(games)
}

// if lower, + 0
// if higher, + difference
// if 21, - what you lost out on

/*
state is just the current-highest
actions are: reroll, hold
since just 2, it can be pos/neg
*/
