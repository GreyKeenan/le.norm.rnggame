package main

import (
	"os"
	"fmt"
	"bufio"
)

type Actor interface {
	Act(state float64) (error, bool)
	Fin(total float64)
}

// ==========

type cli_actor struct { rdr *bufio.Reader }

func Cli_actor() Actor { return cli_actor{bufio.NewReader(os.Stdin)} }

func (self cli_actor) Act(state float64) (error, bool) {

	fmt.Printf("you have %v. Reroll? (y)\n", state)

	var text string
	var e error
	text, e = self.rdr.ReadString('\n')
	if (e != nil) { return e, false }


	if (text[0] != 'y') { return nil, true }

	return nil, false
}
func (self cli_actor) Fin(state float64) {
	fmt.Printf("you won %v!\n", state)
}
