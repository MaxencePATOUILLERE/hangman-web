package main

import (
	"encoding/json"
	"github.com/gorilla/sessions"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	webPageMapSolo = map[int]GameWebPageSolo{}
	connListSolo   = map[int]*websocket.Conn{}
)

type GameWebPageSolo struct {
	Word     string
	Used     string
	Attempts int
	Finish   bool
}

func handleWebSocketSolo(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	connListSolo[session.Values["uid"].(int)] = conn
	for {
		// Read message from browser
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		// Print the message to the console
		log.Print(session.Values["uid"], string(msg))
		if string(msg) == ": AskInfoSend" {
			genWebData(session)
			broadCastStateSolo()
		}else if string(msg) != "" {
			temp := dictPlayer[session.Values["uid"].(int)]
			if len(string(msg)) == 1 {
				if !isGood(temp.userData.ToFind, rune(string(msg)[0])) {
					temp.userData.Attempts++
				}
			} else if string(msg) == temp.userData.ToFind {
				temp.userData.Word = string(msg)
			}
			temp.userData = trys(temp.userData, rune(string(msg)[0]))
			dictPlayer[session.Values["uid"].(int)] = temp
			genWebData(session)
			broadCastStateSolo()
		}
	}
}

func genWebData(session *sessions.Session) {
	for uid, _ := range connListSolo {
		datas := dictPlayer[session.Values["uid"].(int)].userData
		webPageMapSolo[uid] = GameWebPageSolo{Used: string(datas.Used), Word: datas.Word, Attempts: datas.Attempts, Finish: finish(datas)}
	}
}

func broadCastStateSolo() {
	for uid, con := range connListSolo {
		datas, err := json.Marshal(webPageMapSolo[uid])
		log.Println(uid, string(datas))
		if err != nil {
			log.Println(err)
		}
		err = con.WriteMessage(1, datas)
		if err != nil {
			log.Println(err)
		}
	}
}

/*func gamePage(w http.ResponseWriter, r *http.Request, UData *UserData) {
	session, _ := store.Get(r, "cookie-name")
	tmpl := template.Must(template.ParseFiles("assets/web/soloTempl.html"))
	if r.FormValue("letter") != "" {
		if len(r.FormValue("letter")) == 1 && isGood(UData.userData.ToFind, rune(r.FormValue("letter")[0])) {
			UData.userData = trys(UData.userData, rune(r.FormValue("letter")[0]))
		} else if len(r.FormValue("letter")) > 1 {
			UData.userData = guessWord(UData.userData, r.FormValue("letter"))
		} else {
			UData.userData.Attempts++
			UData.userData.HangManState = printHangMan(UData.userData.Attempts)
		}
	}
	if UData.inGame {
		tmpl.Execute(w, UData.userData)
	} else {
		fmt.Fprintf(w, "You are not in the game")
	}
}*/
