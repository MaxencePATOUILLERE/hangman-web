package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"net/http"
	"text/template"
)

type formDisp struct {
	PlayerDisp []string
	PlayerList []string
}

type WebPage struct {
	PlayerName string
	GameUData  UserData
	Form       formDisp
	PlayerList []string
	Admin      bool
}

type UserData struct {
	uid      interface{}
	userData HangManData
	userName string
	admin    bool
	finish   bool
}

type MultiplayerWebPage struct {
	Users    []UserData
	UserList [][]string
	Started  bool
	Mode     string
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	usernames       = []string{"Master Chief", "Link", "Steve", "Kratos"}
	multiplayerData = MultiplayerWebPage{}
	playerList      []string
	dictPlayer      = map[interface{}]UserData{}
	key             = []byte("clef super secrete")
	store           = sessions.NewCookieStore(key)
	data            = WebPage{}
)

func displayPage(w http.ResponseWriter, r *http.Request, data WebPage) {
	tmpl := template.Must(template.ParseFiles("assets/web/game.html"))
	if !multiplayerData.Started {
		formFunc(*r, w, data)
		data.GameUData = dictPlayer[data.GameUData.uid]
		tmpl.Execute(w, data)
	} else {
		data.Admin = false
		data.GameUData = dictPlayer[data.GameUData.uid]
		switch multiplayerData.Mode {
		case "versus":
			gamePage(w, r, &data.GameUData)
			versusGame(w, r, &multiplayerData)
		case "coop":
			fmt.Println("Coop")
		case "solo":
			gamePage(w, r, &data.GameUData)
		}
		dictPlayer[data.GameUData.uid] = data.GameUData
	}
}

func main() {
	http.HandleFunc("/", onConnect)
	http.ListenAndServe(":8080", nil)
}
