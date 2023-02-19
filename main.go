package main

import "oop/dfa"

// import (
// 	"fmt"
// )

// type DFA struct {
// 	numberOfStates int
// 	// alphabet [] string
// 	// transitionsTable []map[string]int
// 	// acceptStates [] int
// 	startState int
// }

// // func New(numberOfStates int, startState int) DFA {
// // 	machine := DFA{numberOfStates, startState}
// // 	return machine
// // }

// func (machine DFA) run() {

// }

// func (machine DFA) print() {
// 	fmt.Println(machine.numberOfStates)
// }

func main() {
	var machine = dfa.DFA()
	machine.print()
}
