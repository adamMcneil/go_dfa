package dfa

import (
	"fmt"
)

type DFA struct {
	numberOfStates int
	// alphabet [] string
	// transitionsTable []map[string]int
	// acceptStates [] int
}

func New(numberOfStates int) DFA {
	machine := DFA{numberOfStates}
	return machine
}

func (machine DFA) run() {

}

func (machine DFA) print() {
	fmt.Println(machine.numberOfStates)
}
