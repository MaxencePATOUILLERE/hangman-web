package main

import (
	"html/template"
	"net/http"
)

func main() {
	menuScreen()
}

func menuScreen() {
	tmpl := template.Must(template.ParseFiles("assets/web/index.html"))
	GameData := setup("./assets/words/words.txt", "./assets/standard.txt")
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			if r.FormValue("letter") != "" {
				if isGood(GameData.ToFind, rune(r.FormValue("letter")[0])) {
					GameData = trys(GameData, rune(r.FormValue("letter")[0]))
				} else {
					GameData.Attempts++
					GameData.HangManState = printHangMan(GameData.Attempts)
				}
			}
			tmpl.Execute(w, GameData)
		})
	http.ListenAndServe(":80", nil)
}
