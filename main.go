package main

//Apres socket
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
	turn     bool
	inGame   bool
}

type MultiplayerData struct {
	Users    []UserData
	UserList []string
	GameData HangManData
	Started  bool
	Mode     string
	Turn     string
}

var (
	usernames       = []string{"Master Chief", "Steve", "Kratos"}
	multiplayerData = MultiplayerData{}
	playerList      []string
	dictPlayer      = map[interface{}]UserData{}
	key             = []byte("clef super secrete")
	store           = sessions.NewCookieStore(key)
	dictTurn        = map[int]bool{}
	uidList         = []int{}
)

func displayPage(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "cookie-name")
	UData := dictPlayer[session.Values["uid"]]
	data := WebPage{
		PlayerName: UData.userName,
		GameUData:  UData,
		PlayerList: playerList,
	}
	if !multiplayerData.Started {
		tmpl := template.Must(template.ParseFiles("assets/web/game.html"))
		formFunc(*r, w, data)
		data.GameUData = dictPlayer[data.GameUData.uid]
		tmpl.Execute(w, data)
	} else {
		data.GameUData = dictPlayer[data.GameUData.uid]
		switch multiplayerData.Mode {
		case "versus":
			versus(w, r)
		case "coop":
			fmt.Println("Coop")
		case "solo":
			gamePage(w, r, &data.GameUData)
		}
		dictPlayer[data.GameUData.uid] = data.GameUData
	}
}

func main() {
	//hub := newHub() //Creation et
	//go hub.run()    //lancement du hub
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/", onConnect)
	/*http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})*/
	http.ListenAndServe(":8080", nil)
}
