package game

import (
	"math/rand"

	"os"
	"fmt"
	"bufio"
	"io"
	"time"
)

const minimum int = 1; //inclusive
const maximum int = 21; //inclusive

func query(input *bufio.Reader, output io.Writer) bool {
	var text string
	for {
		fmt.Fprintf(output, "Roll again? (y/n) : ")
		text, _ = input.ReadString('\n')

		switch text[0] {
			case 'y', 'Y':
				return false
			case 'n', 'N':
				return true
			default:
				fmt.Fprintf(output, "[unrecognized response] ")
		}
	}
}

func generate(output io.Writer, r *rand.Rand) int {

	var n int
	for i := 0; i < 75; i++ {
		n = minimum + r.Intn(maximum - minimum + 1)
		fmt.Fprintf(output, "\r[%2d] ", n)
		time.Sleep(10 * time.Millisecond)
	}
	return n
}

func Play(input *bufio.Reader, output io.Writer, rzer *rand.Rand) bool {
	/*
	assumes maximum - minumum > 0

	generate ranom number [1, 21]
	if n == 21 the game ends and they lose whatever they have
	otherwise they can choose to reroll or not
	cache out:
		"win" the highest number they've seen so far, other than 21
	keep going:
		repeat
	*/

	var n int
	var highest int = 0
	for {
		n = generate(output, rzer)
		if n == maximum {
			fmt.Fprintf(output, "Unlucky!\n")
			if (highest != 0) {
				fmt.Fprintf(output, "\nYou lost $%d (fool)\n", highest)
			}
			return true
		}
		if n > highest { highest = n }
		if query(input, output) { break }
	}

	if highest == maximum - 1 {
		fmt.Fprintf(output, "\nYou got $%d!", highest)
	} else {
		fmt.Fprintf(output, "\nYou got $%d (coward)\n", highest)
	}
	return false
}

func Play_cli(r *rand.Rand) bool {
	return Play(bufio.NewReader(os.Stdin), os.Stdout, r)
}
