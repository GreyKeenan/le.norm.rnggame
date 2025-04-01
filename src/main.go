package main

import (

	//"github.com/greykeenan/le.norm.rnggame/src/game"
	"github.com/greykeenan/le.norm.rnggame/src/toddler"

	"math/rand"
	"time"
	"fmt"
	"os"
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

	// seed with immediate '21' loss: 1743473300

	var src rand.Source = rand.NewSource(seed)
	var rzer *rand.Rand = rand.New(src)

	/*

	print("\033[H\033[2J")
	fmt.Printf("\nseed: %d\n\n", seed)
	game.Play_cli(rzer)
	fmt.Println()

	*/

	var counts [toddler.Possibilities]int
	var totals [toddler.Possibilities]int
	var averages [toddler.Possibilities]float64

	toddler.Train(&counts, &totals, rzer, 100000)
	toddler.Solidify(&counts, &totals, &averages)

	var score float64 = toddler.Perform(&averages, rzer, 10000)

	fmt.Printf("\nseed: %d\n\n", seed)
	fmt.Println(counts)
	fmt.Println(totals)
	fmt.Println(averages)
	fmt.Println()
	fmt.Printf("Score: %v\n", score)
	fmt.Println()

	return
}
