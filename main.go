package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	menuScreen()
}

var (
	GameData = setup("./assets/words/words.txt", "./assets/standard.txt")
)

func menuScreen() {
	fs := http.FileServer(http.Dir("assets/web/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/", gamePage)
	fmt.Println("(http://localhost:80) - Server started on port 80")
	http.ListenAndServe(":80", nil)
}

func gamePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("assets/web/index.html"))
	if r.FormValue("letter") != "" {
		if isGood(GameData.ToFind, rune(r.FormValue("letter")[0])) {
			GameData = trys(GameData, rune(r.FormValue("letter")[0]))
		} else {
			GameData.Attempts++
			GameData.HangManState = printHangMan(GameData.Attempts)
		}
	}
	tmpl.Execute(w, GameData)
}
