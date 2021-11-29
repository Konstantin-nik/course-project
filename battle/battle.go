package main

// type Battle interface {
// 	DoDamage(b Battle, d *Damage) (Status, string)
// 	GetDamage(d *Damage) (Status, string)
// 	PrintStatus()
// 	UpdateStatus() int
// 	isAlive() bool
// }

type Arena interface {
	UpdateStatus()
	GetResult() (a Battle)
}

type BattlePair struct {
	b1, b2 Battle
	status bool
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

type Battle interface {
	IsAlive() bool
	GetDamage(d float64)
	DoDamage(i interface{ GetDamage(d float64) })
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

/*
type Damage struct {
	Value float64
	Type  string
	Power int
}

type Warrior struct {
	Health float64
	Armor  float64
	Status Status
}

type Mage struct {
	Health float64
	Mana   float64
	Status Status
}

func (m *Mage) DoDamage(b Battle, d *Damage) (Status, string) {
	m.UpdateStatus()
	if !m.isAlive() {
		m.Status.Name = "dead"
		m.Status.Value = 0
		return m.Status, "Mage"
	}
	fmt.Print("Mage do ")
	return b.GetDamage(d)
}

func (w *Warrior) DoDamage(b Battle, d *Damage) (Status, string) {
	w.UpdateStatus()
	if !w.isAlive() {
		w.Status.Name = "dead"
		w.Status.Value = 0
		return w.Status, "Warrior"
	}
	fmt.Print("Warrior do ")
	return b.GetDamage(d)
}

func (m *Mage) GetDamage(d *Damage) (Status, string) {
	switch d.Type {
	case "fire":
		m.Health -= d.Value
		fmt.Printf("to Mage %g fire damage\n", d.Value)
		if m.Status.Name == "fire" {
			m.Status.Value += d.Power
		} else if m.Status.Name == "freeze" {
			differ := m.Status.Value - d.Power
			if differ == 0 {
				m.Health -= float64(d.Power)
				fmt.Printf("Mage get %d damage due to collisions of ice and fire", d.Power)
				m.Status.Name = "none"
				m.Status.Value = 0
			} else if differ > 0 {
				m.Health -= float64(d.Power)
				fmt.Printf("Mage get %d damage due to collisions of ice and fire", d.Power)
				m.Status.Value = differ
			} else if differ < 0 {
				m.Health -= float64(m.Status.Value)
				fmt.Printf("Mage get %d damage due to collisions of ice and fire", m.Status.Value)
				m.Status.Name = "fire"
				m.Status.Value = -differ
			}
		} else {
			m.Status.Name = "fire"
			m.Status.Value = d.Power
		}
	case "freeze":
		m.Health -= d.Value
		fmt.Printf("to Mage %g ice damage\n", d.Value)
		if m.Status.Name == "freeze" {
			m.Status.Value += d.Power
		} else if m.Status.Name == "fire" {
			differ := m.Status.Value - d.Power
			if differ == 0 {
				m.Health -= float64(d.Power)
				fmt.Printf("Mage get %d damage due to collisions of ice and fire", d.Power)
				m.Status.Name = "none"
				m.Status.Value = 0
			} else if differ > 0 {
				m.Health -= float64(d.Power)
				fmt.Printf("Mage get %d damage due to collisions of ice and fire", d.Power)
				m.Status.Value = differ
			} else if differ < 0 {
				m.Health -= float64(m.Status.Value)
				fmt.Printf("Mage get %d damage due to collisions of ice and fire", m.Status.Value)
				m.Status.Name = "freeze"
				m.Status.Value = -differ
			}
		} else {
			m.Status.Name = "freeze"
			m.Status.Value = d.Power
		}
	default:
		fmt.Printf("to Mage %g damage\n", d.Value)
		m.Health -= d.Value
	}
	if !m.isAlive() {
		m.Status.Name = "dead"
		m.Status.Value = 0
	}
	m.PrintStatus()
	return m.Status, "Mage"
}

func (w *Warrior) GetDamage(d *Damage) (Status, string) {
	switch d.Type {
	case "fire":
		w.Health -= d.Value - w.Armor
		fmt.Printf("to Warrior %g fire damage\n", d.Value-w.Armor)
		if w.Status.Name == "fire" {
			w.Status.Value += d.Power
		} else if w.Status.Name == "freeze" {
			differ := w.Status.Value - d.Power
			if differ == 0 {
				w.Health -= float64(d.Power)
				fmt.Printf("Warrior get %d damage due to collisions of ice and fire", d.Power)
				w.Status.Name = "none"
				w.Status.Value = 0
			} else if differ > 0 {
				w.Health -= float64(d.Power)
				fmt.Printf("Warrior get %d damage due to collisions of ice and fire", d.Power)
				w.Status.Value = differ
			} else if differ < 0 {
				w.Health -= float64(w.Status.Value)
				fmt.Printf("Warrior get %d damage due to collisions of ice and fire", w.Status.Value)
				w.Status.Name = "fire"
				w.Status.Value = -differ
			}
		} else {
			w.Status.Name = "fire"
			w.Status.Value = d.Power
		}
	case "freeze":
		w.Health -= d.Value - w.Armor
		fmt.Printf("to Warrior %g ice damage\n", d.Value-w.Armor)
		if w.Status.Name == "freeze" {
			w.Status.Value += d.Power
		} else if w.Status.Name == "fire" {
			differ := w.Status.Value - d.Power
			if differ == 0 {
				w.Health -= float64(d.Power)
				fmt.Printf("Warrior get %d damage due to collisions of ice and fire", d.Power)
				w.Status.Name = "none"
				w.Status.Value = 0
			} else if differ > 0 {
				w.Health -= float64(d.Power)
				fmt.Printf("Warrior get %d damage due to collisions of ice and fire", d.Power)
				w.Status.Value = differ
			} else if differ < 0 {
				w.Health -= float64(w.Status.Value)
				fmt.Printf("Warrior get %d damage due to collisions of ice and fire", w.Status.Value)
				w.Status.Name = "freeze"
				w.Status.Value = -differ
			}
		} else {
			w.Status.Name = "freeze"
			w.Status.Value = d.Power
		}
	case "piercing":
		w.Health -= d.Value
		fmt.Printf("to Warrior %g piercing damage\n", d.Value)
	default:
		w.Health -= d.Value - w.Armor
		fmt.Printf("to Warrior %g damage\n", d.Value-w.Armor)
	}
	if !w.isAlive() {
		w.Status.Name = "dead"
		w.Status.Value = 0
	}
	w.PrintStatus()
	return w.Status, "Warrior"
}

func (m *Mage) PrintStatus() {
	fmt.Printf("Mage Status: (Name: Mage, Health: %g, Mana: %g, Status: %s)\n", m.Health, m.Mana, m.Status.Name)
}

func (w *Warrior) PrintStatus() {
	fmt.Printf("Warrior Status: (Name: Warrior, Health: %g, Armor: %g, Status: %s)\n", w.Health, w.Armor, w.Status.Name)
}

func (m *Mage) UpdateStatus() int {
	switch m.Status.Name {
	case "fire":
		m.Health -= float64(m.Status.Value)
		fmt.Printf("Mage get %d damage from fire\n", m.Status.Value)
		m.Status.Value -= 1
		m.PrintStatus()
		return 0
	case "freeze":
		m.Status.Value -= 1
		return 1
	default:
		return 0
	}
}

func (w *Warrior) UpdateStatus() int {
	switch w.Status.Name {
	case "fire":
		w.Health -= float64(w.Status.Value)
		fmt.Printf("Warrior get %d damage from fire\n", w.Status.Value)
		w.Status.Value -= 1
		w.PrintStatus()
		return 0
	case "freeze":
		w.Status.Value -= 1
		return 1
	default:
		return 0
	}
}

func (m *Mage) isAlive() bool {
	return m.Health > 0
}

func (w *Warrior) isAlive() bool {
	return w.Health > 0
}

func main() {
	fmt.Println("Welcome to the arena!")
	fmt.Println("Today we have two players: Mage and Warrior!")
	fmt.Println("Lets have a look to our Players: ")
	var p1 Battle = &Mage{Health: 100.0, Mana: 0.0, Status: Status{Name: "Alive", Value: 0}}
	fmt.Println("First Player: ")
	p1.PrintStatus()
	var p2 Battle = &Warrior{Health: 100.0, Armor: 5.0, Status: Status{Name: "Alive", Value: 0}}
	fmt.Println("Second Player: ")
	p2.PrintStatus()
	fmt.Println("\nFight!")
	var s1 Status = Status{Name: "None", Value: 0}
	var s2 Status = Status{Name: "None", Value: 0}
	var name1, name2 string
	for s1.Name != "dead" && s2.Name != "dead" {
		s1, name1 = p1.DoDamage(p2, &Damage{Value: 10.0, Type: "fire", Power: 6})
		fmt.Println("***********************************")
		s2, name2 = p2.DoDamage(p1, &Damage{Value: 20.0, Type: "normal", Power: 0})
		fmt.Println("***********************************")
	}

	if s1.Name == "dead" && s2.Name == "dead" {
		fmt.Println("Both Players died")
	} else if s1.Name == "dead" {
		if name1 == "Warrior" {
			fmt.Printf("Mage won\n")
		} else if name1 == "Mage" {
			fmt.Printf("Warrior won\n")
		}
	} else if s2.Name == "dead" {
		if name2 == "Warrior" {
			fmt.Printf("Mage won\n")
		} else if name2 == "Mage" {
			fmt.Printf("Warrior won\n")
		}
	}
	fmt.Scan()
}
*/
