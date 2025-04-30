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

/*
Initially, I thought the ideal cap would be `9.5`.
At that point, 50% of the next values are wins.
However, that means you're getting `9.5+`, 50% of the time,
which is only like `4.25+` on average.

And, yes, that's what I observe with a cap of `9.5`.
Testing different values, around 8 is optimal.

With 8.5, you're getting `8.5+`, 80-something percent of the time.
With 7.5, you're getting `7.5+`, 90-something percent of the time.
So, it makes sense that its somewhere between there.
*/
