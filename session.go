package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Value: "Player " + string(len(playerList)), Expires: time.Now().AddDate(0, 0, 1)}
	/*http.Error(w, "Game Full or already started", 403)
	return*/
	r.AddCookie(&c)
}

func logout(w http.ResponseWriter, r *http.Request, pUDATA UserData) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func onConnect(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth && len(playerList) < 3 {
		session.Values["authenticated"] = true
		rand.Seed(time.Now().UnixNano())
		uid := rand.Intn(100000)
		session.Values["uid"] = uid
		session.Save(r, w)
		uidList = append(uidList, uid)
		dictTurn[uid] = false
		dictPlayer[session.Values["uid"]] = setupGameData(session)
		multiplayerData.Users = append(multiplayerData.Users, dictPlayer[session.Values["uid"]])
		multiplayerData.UserList = playerList
		playerList = append(playerList, usernames[len(playerList)])
	} else if len(playerList) == 3 && !auth {
		fmt.Fprintf(w, "Game Full")
		return
	}
	displayPage(w, r)
}
