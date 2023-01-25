package main

import (
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"net/http"
)

type HangManData struct {
	Word     string
	ToFind   string
	Attempts int
	Used     []rune
}

type UserData struct {
	userName string
	userData HangManData
	admin    bool
	finish   bool
	turn     bool
	inGame   bool
}

var (
	dictPlayer = map[int]UserData{}
	usernames  = []string{"Master Chief", "Steve", "Kratos"}
	playerList []string
	key        = []byte("clef super secrete")
	store      = sessions.NewCookieStore(key)
	upgrader   = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func reset() {
	playerList = []string{}
	dictPlayer = map[int]UserData{}

	GData = HangManData{}
	multiplayerData = MultiplayerData{}
	webPageMapMulti = map[int]GameWebPageMulti{}
	connListMulti = map[int]*websocket.Conn{}

	webPageMapSolo = map[int]GameWebPageSolo{}
	connListSolo = map[int]*websocket.Conn{}

	connListHub = map[int]*websocket.Conn{}
	webPageMapHub = map[int]HubWebPage{}
}

func main() {
	fi := http.FileServer(http.Dir("./assets/web/"))
	http.Handle("/web/", http.StripPrefix("/web/", fi))

	http.HandleFunc("/multi/ws", handleWebSocketMulti)
	http.HandleFunc("/multi", func(w http.ResponseWriter, r *http.Request) {
		//onConnect(w, r)
		http.ServeFile(w, r, "./assets/web/multiplayer.html")
	})
	http.HandleFunc("/solo/ws", handleWebSocketSolo)
	http.HandleFunc("/solo", func(w http.ResponseWriter, r *http.Request) {
		//onConnect(w, r)
		http.ServeFile(w, r, "./assets/web/solo.html")
	})
	http.HandleFunc("/ws", handleWebSocketHub)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		onConnect(w, r)
		http.ServeFile(w, r, "./assets/web/hub.html")
	})
	http.ListenAndServe(":8080", nil)
}
