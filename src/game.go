package main

import (
	"math/rand"
)


func play(rng *rand.Rand, actor Actor) (error, float64) {

	var e error

	var state float64 = 0
	var response bool

	for {
		e, response = actor.Act(state)
		if (e != nil) { return e, state }
		if (response) { break }

		state += rng.NormFloat64() + 0.5
	
		if state >= 10 {
			state = 0
			break
		}
	}

	actor.Fin(state)

	return nil, 0;
}

