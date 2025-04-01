package main

import (

	"github.com/greykeenan/le.norm.rnggame/src/game"

	"math/rand"
	"time"
	"fmt"
)

func main() {

	var t int64 = time.Now().Unix()
	var src rand.Source = rand.NewSource(t)
	var rzer *rand.Rand = rand.New(src)

	print("\033[H\033[2J")
	fmt.Printf("\nseed: %d\n\n", t)
	game.Play_cli(rzer)
	fmt.Println()

	return
}
