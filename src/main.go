package main

import (
	"math/rand"
	"os"
	"fmt"
	"time"
	"strconv"
)

func main() {

	var e error

	var cap_at float64
	if len(os.Args) > 1 {
		cap_at, e = strconv.ParseFloat(os.Args[1], 64)
		if (e != nil) { panic(e) }
	} else {
		cap_at = 8
	}
	fmt.Printf("cap at: %v\n", cap_at)

	var seed int64
	if len(os.Args) > 2 {
		seed, e = strconv.ParseInt(os.Args[2], 10, 63)
		if (e != nil) {
			seed = time.Now().Unix()
		}
	} else {
		seed = time.Now().Unix()
	}
	fmt.Printf("seed: %v\n", seed);
	var rng *rand.Rand = rand.New(rand.NewSource(seed))

	/*
	var winnings float64
	e, winnings = play(rng, Cli_actor())
	if (e != nil) { panic(e) }

	fmt.Printf("You won %v.\n", winnings)
	*/

	var avg_winnings float64
	e, avg_winnings = play_average(rng, Static_actor(cap_at), 10000)

	if (e != nil) { panic(e) }

	fmt.Printf("won %v on average\n", avg_winnings)
}
