package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	LevelForm        string
	LevelFormDisplay bool
	GameDisp         GameDisplayForm
}

type GameDisplayForm struct {
	WordDisp     string
	HangManState string
	UsedLetters  []rune
	InputArea    string
}

func main() {
	menuScreen()
}

func menuScreen() {

	tmpl := template.Must(template.ParseFiles("assets/web/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
		GameData := set(r.FormValue("difficulty")+".txt", "")
		template.Must(template.New("tpl").Parse("<h1>{{.Word}}</h1><form method=\"POST\"><label>Enter a letter :</label><br/><input type=\"text\" name=\"letter\"><br/><input type=\"submit\">")).Execute(w, GameData)
	})
	http.ListenAndServe(":80", nil)
}

func gameDisplay(data HangManData, w http.ResponseWriter, r *http.Request) {
	out := r.FormValue("letter")
	if len(out) > 0 {
		fmt.Println(out)
		trys(data, rune(out[0]))
	}

}
