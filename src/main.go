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

	var seed int64
	if len(os.Args) > 1 {
		seed, e = strconv.ParseInt(os.Args[1], 10, 63)
		if (e != nil) {
			seed = time.Now().Unix()
		}
	} else {
		seed = time.Now().Unix()
	}
	var rng *rand.Rand = rand.New(rand.NewSource(seed))

	/*
	var winnings float64
	e, winnings = play(rng, Cli_actor())
	if (e != nil) { panic(e) }

	fmt.Printf("You won %v.\n", winnings)
	*/

	var avg_winnings float64
	e, avg_winnings = play_average(rng, Static_actor(8), 10000)

	if (e != nil) { panic(e) }

	fmt.Printf("won %v on average\n", avg_winnings)
}
