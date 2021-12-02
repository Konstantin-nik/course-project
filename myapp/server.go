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

func (p *PlayerInfo) GetPlayerInfo(l *PlayerList) {
	p.Players = make([]Player, len(l.Players))
	for i := 0; i < len(l.Players); i++ {
		p.Players[i].Index = i + 1
		p.Players[i].Name = (*l.Players[i]).Name()
		p.Players[i].Health = (*l.Players[i]).Health()
		p.Players[i].Status = (*l.Players[i]).Status()
	}
}

type WinnerInfo struct {
	Winner bool
	Name   string
	Health string
	Status string
}

func (winfo *WinnerInfo) GetWinnerInfo(winner *btl.Player) {
	if winner == nil {
		winfo.Winner = false
	} else {
		winfo.Winner = true
		winfo.Name = (*winner).Name()
		winfo.Health = (*winner).Health()
		winfo.Status = (*winner).Status()
	}
}

func main() {
	p := &PlayerList{}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.Handle("/addplayer", p)

	http.HandleFunc("/info", func(w http.ResponseWriter, _ *http.Request) {
		var data *PlayerInfo = &PlayerInfo{}
		data.GetPlayerInfo(p)
		tmpl, _ := template.ParseFiles("templates/info.html")
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/battle", func(w http.ResponseWriter, _ *http.Request) {
		winner := p.Battle()
		var data *WinnerInfo
		data.GetWinnerInfo(winner)
		tmpl, _ := template.ParseFiles("templates/battle.html")
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "css/styles.css")
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
