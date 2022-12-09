package main

import (
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
	WebPageData := WebPage{}
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		WebPageData.GameUData = setupGameData(session)
		session.Values["authenticated"] = true
		rand.Seed(time.Now().UnixNano())
		session.Values["uid"] = rand.Intn(100000 - 0)
		session.Save(r, w)
		playerList = append(playerList, usernames[len(playerList)])
	} else {
		WebPageData.GameUData = dictPlayer[session.Values["uid"]]
	}
	WebPageData.Admin = WebPageData.GameUData.admin
	WebPageData.PlayerName = WebPageData.GameUData.userName
	WebPageData.PlayerList = playerList
	dictPlayer[session.Values["uid"]] = WebPageData.GameUData
	displayPage(w, r, WebPageData)
}

func isAuthorised(userData UserData) {
}
