package main

import (
	"fmt"
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

func start(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	tmpl := template.Must(template.ParseFiles("assets/web/game.html"))
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		login(w, r)
		setPlayerMultiplayerData(w, r, session)
	} else {
		data.PlayerName = dictPlayer[session.Values["uid"]].userName
	}
	data.PlayerList = playerList
	if !multiplayerData.Started {
		fmt.Println(data.GameUData.userData.Word, data.GameUData.userData.ToFind)
		tmpl.Execute(w, data)
		checkMultiplayerForm(w, r, session)
		dictPlayer[session.Values["uid"]] = data.GameUData
	} else {
		data.Admin = false
		switch multiplayerData.Mode {
		case "versus":
			/*gamePage(w, r, &data.GameUData)
			versusGame(w, r, &multiplayerData)*/
			fmt.Println("Coop")
		case "coop":
			fmt.Println("Coop")
		case "solo":
			gamePage(w, r, &data.GameUData)
		}
		dictPlayer[data.GameUData.uid] = data.GameUData
	}
}

func main() {
	http.HandleFunc("/", start)
	http.ListenAndServe(":8080", nil)
}
