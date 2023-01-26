package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type HubWebPage struct {
	PlayerList []string
	Admin      bool
	Redirect   string
	Type       string
}

type JSONInput struct {
	Mode       string `json:"Mode"`
	Difficulty string `json:"Difficulty"`
}

var (
	connListHub   = map[int]*websocket.Conn{}
	webPageMapHub = map[int]HubWebPage{}
)

func handleWebSocketHub(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	connListHub[session.Values["uid"].(int)] = conn
	for {
		// Read message from browser
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		log.Print(session.Values["uid"], string(msg))
		genHubWebPage()
		broadcastHub()
		//Insert Form Get Request
		dat := JSONInput{}
		println(string(msg))
		json.Unmarshal(msg, &dat)
		mode, difficulty := dat.Mode, dat.Difficulty
		fmt.Println(mode, difficulty)
		//End
		if string(msg) == ": AskInfoSend"{
			genWebData(session)
			temp, _ := json.Marshal(webPageMapHub[session.Values["uid"].(int)])
			conn.WriteMessage(1, temp)
		}else if string(msg) != ""  {
			if mode == "multi" {
				GData = generateGamesDatas(difficulty)
				redirectToUrl("multi")
				conn.Close()
			} else if mode == "solo" {
				temp := dictPlayer[session.Values["uid"].(int)]
				temp.userData = generateGamesDatas(difficulty)
				dictPlayer[session.Values["uid"].(int)] = temp
				redirectToUrl("solo")
				conn.Close()
			}
		}
	}
}

func generateGamesDatas(difficulty string) HangManData {
	wl := ""
	switch difficulty {
	case "easy":
		wl = formatWord(getFileWords("assets/words/words.txt"))
	case "hard":
		wl = formatWord(getFileWords("assets/words/words3.txt"))
	case "medium":
		wl = formatWord(getFileWords("assets/words/words2.txt"))
	}
	HData := reveal(HangManData{Attempts: 0, ToFind: wl, Used: []rune{}, Word: genHidden(wl)})
	return HData
}

func genHubWebPage() {
	for uid, _ := range connListHub {
		webPageMapHub[uid] = HubWebPage{PlayerList: playerList, Admin: dictPlayer[uid].admin, Type: "datas"}
	}
}

func redirectToUrl(redir string) {
	for _, conn := range connListHub {
		datas, _ := json.Marshal(HubWebPage{Redirect: redir, Type: "redirect"})
		conn.WriteMessage(1, datas)
	}
}

func broadcastHub() {
	for uid, conn := range connListHub {
		datas, _ := json.Marshal(webPageMapHub[uid])
		conn.WriteMessage(1, datas)
	}
}
