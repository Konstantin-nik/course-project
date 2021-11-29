package main

import (
	"fmt"
)

var (
	player1 Battle = &Warrior{P: Person{Name: "player1", Health: 200}, Damage: 10,
		Armor: 5, Range: 3, Flee: 0.69}
	player2 Battle = &Warrior{P: Person{Name: "player2", Health: 200}, Damage: 11,
		Armor: 5, Range: 2, Flee: 0.65}
	arena Arena = &BattlePair{b1: player1, b2: player2}
)

func main() {
	fmt.Print("Welcome to the arena")
	arena.StartBattle()
	fmt.Print(arena.GetResult())
}

// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }
