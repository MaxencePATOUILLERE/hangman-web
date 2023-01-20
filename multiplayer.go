package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func versusGame(w http.ResponseWriter, r *http.Request, GData HangManData) {
	session, _ := store.Get(r, "cookie-name")
	tmpl := template.Must(template.ParseFiles("assets/web/versusTempl.html"))
	if r.FormValue("letter") != "" && dictPlayer[session.Values["uid"]].userName == multiplayerData.Turn {
		if len(r.FormValue("letter")) == 1 && isGood(GData.ToFind, rune(r.FormValue("letter")[0])) {
			GData = trys(GData, rune(r.FormValue("letter")[0]))
		} else if len(r.FormValue("letter")) > 1 {
			GData = guessWord(GData, r.FormValue("letter"))
		}
		if finish(GData) {
			fmt.Fprintf(w, multiplayerData.Turn+" is the winner ")
		}
		changePlayerTurn()
		multiplayerData.GameData = GData
	}
	tmpl.Execute(w, GData)
}

func versus(w http.ResponseWriter, r *http.Request) {
	versusGame(w, r, multiplayerData.GameData)
}

func formFunc(r http.Request, w http.ResponseWriter, data WebPage) {
	if r.FormValue("difficulty") != "" && r.FormValue("mode") != "" && !multiplayerData.Started {
		multiplayerData.Started = true
		wl := ""
		switch r.FormValue("difficulty") {
		case "easy":
			wl = formatWord(getFileWords("assets/words/words.txt"))
		case "hard":
			wl = formatWord(getFileWords("assets/words/words3.txt"))
		case "medium":
			wl = formatWord(getFileWords("assets/words/words2.txt"))
		}
		mwl := genHidden(wl)
		multiplayerData.Started = true
		switch r.FormValue("mode") {
		case "versus":
			multiplayerData.Mode = "versus"
			multiplayerData.Turn = multiplayerData.Users[0].userName
			multiplayerData.GameData.ToFind = wl
			multiplayerData.GameData.Word = mwl
			versus(w, &r)
		case "coop":
			multiplayerData.Mode = "coop"
		case "solo":
			data.GameUData.userData.ToFind = wl
			data.GameUData.userData.Word = mwl
			data.GameUData.inGame = true
			multiplayerData.Mode = "solo"
			dictPlayer[data.GameUData.uid] = data.GameUData
			gamePage(w, &r, &data.GameUData)
		}
	}
}

func changePlayerTurn() {
	for i := 0; i < len(multiplayerData.Users); i++ {
		if multiplayerData.Users[i].userName == multiplayerData.Turn {
			if i == len(multiplayerData.Users)-1 {
				fmt.Println("Previous : ", multiplayerData.Users[i].userName, " Current : ", multiplayerData.Users[0].userName)
				multiplayerData.Turn = multiplayerData.Users[0].userName
				fmt.Println(multiplayerData.Turn, "'s turn")
				return
			} else {
				fmt.Println("Previous : ", multiplayerData.Users[i].userName, " Current : ", multiplayerData.Users[i+1].userName)
				multiplayerData.Turn = multiplayerData.Users[i+1].userName
				fmt.Println(multiplayerData.Turn, "'s turn")
				return
			}
		}
	}
}
