package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func setPlayerMultiplayerData(w http.ResponseWriter, r *http.Request, ses *sessions.Session) {
	nUdata := UserData{
		userData: HangManData{UserName: usernames[len(multiplayerData.Users)]},
		userName: usernames[len(multiplayerData.Users)],
	}
	multiplayerData.Started = false
	if len(playerList) == 0 {
		data.Admin = true
		nUdata.admin = true
	}
	checkWL(r)
	data.PlayerName = usernames[len(multiplayerData.Users)]
	multiplayerData.Users = append(multiplayerData.Users, nUdata)
	playerList = append(playerList, data.PlayerName)
	ses.Values["authenticated"] = true
	rand.Seed(time.Now().UnixNano())
	ses.Values["uid"] = rand.Intn(100000 - 0)
	ses.Save(r, w)
	multiplayerData.Users = append(multiplayerData.Users, nUdata)
	nUdata.uid = ses.Values["uid"]
	dictPlayer[nUdata.uid] = nUdata

}

func checkWL(r *http.Request) {
	switch r.FormValue("difficulty") {
	case "easy":
		data.GameUData.userData.ToFind = formatWord(getFileWords("assets/words/words.txt"))
	case "hard":
		data.GameUData.userData.ToFind = formatWord(getFileWords("assets/words/words3.txt"))
	case "medium":
		data.GameUData.userData.ToFind = formatWord(getFileWords("assets/words/words2.txt"))
	}
	data.GameUData.userData.Word = genHidden(data.GameUData.userData.ToFind)
	data.GameUData.userData = reveal(data.GameUData.userData)
	fmt.Println(data.GameUData.userData.Word, data.GameUData.userData.ToFind)
	data.GameUData.userData.Attempts = 0
	multiplayerData.UserList = append(multiplayerData.UserList, []string{data.GameUData.userName, data.GameUData.userData.Word, "0"})
}

func checkMultiplayerForm(w http.ResponseWriter, r *http.Request, ses *sessions.Session) {
	if r.FormValue("difficulty") != "" && r.FormValue("mode") != "" {
		multiplayerData.Started = true
		switch r.FormValue("mode") {
		case "versus":
			multiplayerData.Mode = "versus"
		case "coop":
			multiplayerData.Mode = "coop"
		case "solo":
			multiplayerData.Started = true
			multiplayerData.Mode = "solo"
			dictPlayer[ses.Values["uid"]] = data.GameUData
			w.Write([]byte{})
		}
	}
}

func versusGame(w http.ResponseWriter, r *http.Request, MData *MultiplayerWebPage) {
	tmpl := template.Must(template.ParseFiles("assets/web/versusTempl.html"))
	tmpl.Execute(w, MData)
}
