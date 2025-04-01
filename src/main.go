package main

import (

	"github.com/greykeenan/le.norm.rnggame/src/game"

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

	print("\033[H\033[2J")
	fmt.Printf("\nseed: %d\n\n", seed)
	game.Play_cli(rzer)
	fmt.Println()

	return
}
