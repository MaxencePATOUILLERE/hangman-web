package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	menuScreen()
}

func menuScreen() {
	tmpl := template.Must(template.ParseFiles("assets/web/index.html"))
	GameData := HangManData{}
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			if r.FormValue("difficulty") != "" {
				GameData.Difficulty = r.FormValue("difficulty")
				word := getFileWords(r.FormValue("difficulty") + ".txt")
				fmt.Println(word)
				GameData.ToFind = word
				GameData.Word = genHidden(word)
			}
			if r.FormValue("letter") != "" {
				if isGood(GameData.ToFind, rune(r.FormValue("letter")[0])) {
					GameData = trys(GameData, rune(r.FormValue("letter")[0]))
				} else {
					GameData.Attempts++
				}
			}
			tmpl.Execute(w, GameData)
		})
	http.ListenAndServe(":80", nil)
}
