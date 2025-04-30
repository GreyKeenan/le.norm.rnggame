package main

import (
	"os"
	"fmt"
	"bufio"
)

type Actor interface {
	Act(state float64) (error, bool)
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

// ==========

type static_actor struct { v float64 }
func Static_actor(v float64) Actor { return static_actor{v} }
func (self static_actor) Act(state float64) (error, bool) { return nil, (state >= self.v) }
