package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type MultiplayerData struct {
	Users []int
	Turn  int
}

type GameWebPageMulti struct {
	Word   string
	Used   string
	Turn   bool
	Finish bool
	Type string
	Redirect string
}

var (
	GData           = HangManData{}
	multiplayerData = MultiplayerData{}
	webPageMapMulti = map[int]GameWebPageMulti{}
	connListMulti   = map[int]*websocket.Conn{}
)

func handleWebSocketMulti(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	connListMulti[session.Values["uid"].(int)] = conn
	for {
		// Read message from browser
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// Print the message to the console
		log.Print(session.Values["uid"], string(msg))
		if string(msg) == ": AskInfoSend" {
			genHiddenData()
			broadCastState()
		}else if string(msg) == ": ResetRequest"{
			resetMultiplayer()
		}else if multiplayerData.Turn == session.Values["uid"] && string(msg) != "" {
			GData = trys(GData, rune(string(msg)[0]))
			if string(msg) == GData.ToFind {
				GData.Word = string(msg)
			}
			changePlayerTurn()
			genHiddenData()
			broadCastState()
		}
	}
}

func genHiddenData() {
	for uid, _ := range connListMulti {
		turn := multiplayerData.Turn != uid
		webPageMapMulti[uid] = GameWebPageMulti{Used: string(GData.Used), Word: GData.Word, Turn: turn, Finish: finish(GData), Type: "datas"}
	}
}

func broadCastState() {
	for uid, con := range connListMulti {
		datas, err := json.Marshal(webPageMapMulti[uid])
		println(uid, string(datas))
		if err != nil {
			log.Println("Erreur : ", err)
		}
		log.Println("Good : ", string(datas))
		err = con.WriteMessage(1, datas)
		if err != nil {
			log.Println("Erreur : ", err)
		}
	}
}

func changePlayerTurn() {
 	for i := 0; i < len(multiplayerData.Users); i++ {
		if multiplayerData.Users[i] == multiplayerData.Turn {
			if i == len(multiplayerData.Users)-1 {
				multiplayerData.Turn = multiplayerData.Users[0]
				return
			} else {
				multiplayerData.Turn = multiplayerData.Users[i+1]
				return
			}
		}
	}
}

func resetMultiplayer(){
	redirectToUrl(".", connListMulti)
	GData           = HangManData{}
	webPageMapMulti = map[int]GameWebPageMulti{}
	connListMulti   = map[int]*websocket.Conn{}
}
