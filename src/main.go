package main

import (
	"math/rand"
	"fmt"
	"os"
	"bufio"
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

	var rdr *bufio.Reader = bufio.NewReader(os.Stdin)

	var total float64
	var text string
	var rn float64

	for total < 10 {

		fmt.Printf("you have %v. Reroll? (y)\n", total)
		text, _ = rdr.ReadString('\n')

		if (len(text) < 1) {
			break
		}
		if (text[0] != 'y') {
			break
		}

		rn = rng.NormFloat64() + 0.5

		total += rn

		if total >= 10 {
			total = 0
			break
		}
	}

	fmt.Printf("you won %v.\n", total)
}
