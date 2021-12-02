package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	btl "github.com/Konstantin-nik/battle-simulator"
)

type PlayerList struct {
	Players []*btl.Player
}

func (p *PlayerList) AddPlayer(player *btl.Player) {
	p.Players = append(p.Players, player)
}

func (p *PlayerList) Battle() *btl.Player {
	return btl.CircleBattle(p.Players)
}

func (p *PlayerList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		name         = r.FormValue("name")
		health, _    = strconv.ParseFloat(r.FormValue("health"), 32)
		damage, _    = strconv.ParseFloat(r.FormValue("damage"), 32)
		flatArmor, _ = strconv.ParseFloat(r.FormValue("flatArmor"), 32)
		Range, _     = strconv.ParseFloat(r.FormValue("range"), 32)
		perArmor, _  = strconv.ParseFloat(r.FormValue("perArmor"), 32)
	)
	copies, _ := strconv.ParseInt(r.FormValue("copies"), 10, 64)
	cp := int(copies)
	if copies > 0 {
		for copy := 0; copy < cp; copy++ {
			p.AddPlayer(btl.MakePlayer(name, health, damage, flatArmor, Range, perArmor))
		}
	} else {
		p.AddPlayer(btl.MakePlayer(name, health, damage, flatArmor, Range, perArmor))
	}

	http.ServeFile(w, r, "static/index.html")
}

func (p *PlayerList) String() string {
	var s string = ""
	for i, pl := range p.Players {
		s += fmt.Sprintf("%d Player: \n%s\n\n", i+1, (*pl).String())
	}
	return s
}

type Player struct {
	Index  int
	Name   string
	Health string
	Status string
}
type PlayerInfo struct {
	Players []Player
}

func main() {
	p := &PlayerList{}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.Handle("/addplayer", p)

	http.HandleFunc("/info", func(w http.ResponseWriter, _ *http.Request) {
		data := PlayerInfo{
			Players: make([]Player, len(p.Players)),
		}
		for i := 0; i < len(data.Players); i++ {
			data.Players[i].Index = i + 1
			data.Players[i].Name = (*p.Players[i]).Name()
			data.Players[i].Health = (*p.Players[i]).Health()
			data.Players[i].Status = (*p.Players[i]).Status()
		}

		tmpl, _ := template.ParseFiles("templates/info.html")
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/battle", func(w http.ResponseWriter, _ *http.Request) {
		winner := p.Battle()
		if winner == nil {
			fmt.Fprint(w, "No one survived!")
		} else {
			data := Player{
				Index:  0,
				Name:   (*winner).Name(),
				Health: (*winner).Health(),
				Status: (*winner).Status(),
			}

			tmpl, _ := template.ParseFiles("templates/battle.html")
			tmpl.Execute(w, data)
		}
	})

	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "css/styles.css")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
