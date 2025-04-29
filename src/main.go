package main

import (
	"math/rand"
	"os"
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

	play(rng, Cli_actor())
}
