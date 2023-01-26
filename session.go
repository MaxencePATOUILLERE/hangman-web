package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/gorilla/sessions"
	"math/rand"
	"net/http"
	"time"
)

func onConnect(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	println(isAlreadyAuth(session))
	if session.Values["uid"] != nil {
		haveAlreadyCookie(session)
		return
	}
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth && len(playerList) <= 3 {
		setDatas(session, r, w)
	} else if len(playerList) == 4 && !auth {
		fmt.Fprintf(w, "Game Full")
		return
	}
}

func setDatas(session *sessions.Session, r *http.Request, w http.ResponseWriter) {
	session.Values["authenticated"] = true
	rand.Seed(time.Now().UnixNano())
	uid := rand.Intn(100000)
	session.Values["uid"] = uid
	multiplayerData.Users = append(multiplayerData.Users, uid)
	if len(playerList) == 0 {
		dictPlayer[uid] = UserData{userName: usernames[len(playerList)], admin: true}
	} else {
		dictPlayer[uid] = UserData{userName: usernames[len(playerList)], admin: false}
	}
	multiplayerData.Turn = uid
	session.Save(r, w)
	playerList = append(playerList, usernames[len(playerList)])
}

func haveAlreadyCookie(session *sessions.Session) {
	uid := session.Values["uid"].(int)
	for key, _ := range dictPlayer {
		if uid == key {
			return
		}
	}
	multiplayerData.Users = append(multiplayerData.Users, uid)
	if len(playerList) == 0 {
		dictPlayer[uid] = UserData{userName: usernames[len(playerList)], admin: true}
	} else {
		dictPlayer[uid] = UserData{userName: usernames[len(playerList)], admin: false}
	}
	multiplayerData.Turn = uid
	playerList = append(playerList, usernames[len(playerList)])
}

func isAlreadyAuth(session *sessions.Session) bool {
	if session.Values["uid"] != nil{
		return true
	}
	return false
}

func presentInConnList(uid int, connList map[int]*websocket.Conn) bool {
	for key, _ := range connList {
		if key == uid {
			return true
		}
	}
	return false
}
