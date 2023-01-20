package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func gamePage(w http.ResponseWriter, r *http.Request, UData *UserData) {
	session, _ := store.Get(r, "cookie-name")
	UData.userData.HangManState = printHangMan(UData.userData.Attempts)
	tmpl := template.Must(template.ParseFiles("assets/web/soloTempl.html"))
	if r.FormValue("letter") != "" {
		if len(r.FormValue("letter")) == 1 && isGood(UData.userData.ToFind, rune(r.FormValue("letter")[0])) {
			UData.userData = trys(UData.userData, rune(r.FormValue("letter")[0]))
		} else if len(r.FormValue("letter")) > 1 {
			UData.userData = guessWord(UData.userData, r.FormValue("letter"))
		} else {
			UData.userData.Attempts++
			UData.userData.HangManState = printHangMan(UData.userData.Attempts)
		}
		dictPlayer[session.Values["uid"]] = *UData
	}
	if UData.inGame {
		tmpl.Execute(w, UData.userData)
	} else {
		fmt.Fprintf(w, "You are not in the game")
	}
}
