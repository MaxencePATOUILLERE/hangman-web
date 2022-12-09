package main

import (
	"net/http"
	"text/template"
)

func gamePage(w http.ResponseWriter, r *http.Request, UData *UserData) {
	UData.userData.HangManState = printHangMan(UData.userData.Attempts)
	tmpl := template.Must(template.ParseFiles("assets/web/soloTempl.html"))
	if r.FormValue("letter") != "" {
		if isGood(UData.userData.ToFind, rune(r.FormValue("letter")[0])) {
			UData.userData = trys(UData.userData, rune(r.FormValue("letter")[0]))
		} else {
			UData.userData.Attempts++
			UData.userData.HangManState = printHangMan(UData.userData.Attempts)
		}
	}
	tmpl.Execute(w, UData.userData)
}
