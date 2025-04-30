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

	return nil, state;
}

func play_average(rng *rand.Rand, actor Actor, iterations int) (error, float64) {
	
	var e error
	var total float64
	var thisone float64

	var i int
	for i = 0; i < iterations; i += 1 {
		e, thisone = play(rng, actor)
		if (e != nil) { return e, 0 }

		total += thisone
	}

	return nil, total / float64(iterations)
}

