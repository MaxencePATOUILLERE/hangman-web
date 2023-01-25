package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type MultiplayerData struct {
	Users []int
	Turn  int
}

type GameWebPageMulti struct {
	Word string
	Used []rune
	Turn bool
}

var (
	GData           = HangManData{Attempts: 0, ToFind: "test", Word: "____", Used: []rune{}}
	multiplayerData = MultiplayerData{}
	webPageMapMulti = map[int]GameWebPageMulti{}
	gameData        = HangManData{}
	connListMulti   = map[int]*websocket.Conn{}
)

func handleWebSocketMulti(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	connListMulti[session.Values["uid"].(int)] = conn
	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Print the message to the console
		log.Print(session.Values["uid"], string(msg))

		if multiplayerData.Turn == session.Values["uid"] && string(msg) != "" {
			GData = trys(GData, rune(string(msg)[0]))
			genHiddenData()
			changePlayerTurn()
			broadCastState(msgType)

		}
	}
}

func genHiddenData() {
	for uid, _ := range connListMulti {
		turn := multiplayerData.Turn != uid
		webPageMapMulti[uid] = GameWebPageMulti{Used: GData.Used, Word: GData.Word, Turn: turn}
	}
}

func broadCastState(msgT int) {
	for uid, con := range connListMulti {
		datas, err := json.Marshal(webPageMapMulti[uid])
		println(uid, string(datas))
		if err != nil {
			log.Println("Erreur : ", err)
		}
		log.Println("Good : ", string(datas))
		err = con.WriteMessage(msgT, datas)
		if err != nil {
			log.Println("Erreur : ", err)
		}
	}
}

func changePlayerTurn() {
	fmt.Println(multiplayerData.Turn)
	for i := 0; i < len(multiplayerData.Users); i++ {
		if multiplayerData.Users[i] == multiplayerData.Turn {
			if i == len(multiplayerData.Users)-1 {
				fmt.Println("Previous : ", dictPlayer[multiplayerData.Users[i]].userName, " Current : ", dictPlayer[multiplayerData.Users[0]].userName)
				multiplayerData.Turn = multiplayerData.Users[0]
				fmt.Println(dictPlayer[multiplayerData.Turn].userName, "'s turn")
				return
			} else {
				fmt.Println("Previous : ", dictPlayer[multiplayerData.Users[i]].userName, " Current : ", dictPlayer[multiplayerData.Users[i+1]].userName)
				multiplayerData.Turn = multiplayerData.Users[i+1]
				fmt.Println(dictPlayer[multiplayerData.Turn].userName, "'s turn")
				return
			}
		}
	}
}
