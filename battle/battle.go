package main

import (
	"fmt"
)

// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

var (
	player1 Battle = &Warrior{P: Person{Name: "player1", Health: 200}, Damage: 10,
		Armor: 5, Range: 3, Flee: 0.69}
	player2 Battle = &Warrior{P: Person{Name: "player2", Health: 200}, Damage: 11,
		Armor: 5, Range: 2, Flee: 0.65}
	arena Arena = &BattlePair{b1: player1, b2: player2}
)

func main() {
	fmt.Println("Welcome to the arena!")
	arena.StartBattle()
	fmt.Print(arena.GetResult())
}

// Arena interface used for organising battles with multiple players.
//
// List of acceptable structures:
// 		type BattlePair struct {}
// .
type Arena interface {
	StartBattle()
	UpdateStatus()
	GetResult() (a Battle)
}

type BattlePair struct {
	b1, b2 Battle
	status bool
}

func (bp *BattlePair) StartBattle() {
	bp.UpdateStatus()
	for bp.status {
		bp.b1.DoDamage(bp.b2)
		bp.b2.DoDamage(bp.b1)
		bp.UpdateStatus()
	}
}

func (bp *BattlePair) UpdateStatus() {
	bp.status = bp.b1.IsAlive() && bp.b2.IsAlive()
}

func (bp *BattlePair) GetResult() (a Battle) {
	if bp.b1.IsAlive() {
		return bp.b1
	} else if bp.b2.IsAlive() {
		return bp.b2
	} else {
		return nil
	}
}

// Battle interface used for battle skills.
//
// List of acceptable structures:
// 		type Warrior struct {}
// .
type Battle interface {
	IsAlive() bool
	GetDamage(d float64)
	DoDamage(i interface{ GetDamage(d float64) })
	String() string
}

type Status struct {
	Name  string
	Value []int
}

type Person struct {
	Name   string
	Health float64
	Stat   Status
}

func (p *Person) IsAlive() bool {
	return p.Health > 0
}

func (p *Person) GetDamage(d float64) {
	if d > 0 {
		p.Health -= d
	}
}

func (p *Person) String() string {
	return fmt.Sprintf("Status:\n\tName: %s\n\tHealth: %.2f", p.Name, p.Health)
}

type Warrior struct {
	P      Person
	Damage float64
	Armor  float64
	Range  float64
	Flee   float64
}

func (w *Warrior) IsAlive() bool {
	return w.P.IsAlive()
}

func (w *Warrior) GetDamage(d float64) {
	w.P.GetDamage(d - w.Armor)
}

func (w *Warrior) DoDamage(i interface{ GetDamage(d float64) }) {
	i.GetDamage((w.Damage + w.Range) * w.Flee)
}

func (w *Warrior) String() string {
	return w.P.String()
}
