package main

import (
	"net/http"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{Value: "Player " + string(len(playerList)), Expires: time.Now().AddDate(0, 0, 1)}
	if len(playerList) >= 4 || multiplayerData.Started {
		http.Error(w, "Game Full or already started", 403)
		return
	}
	r.AddCookie(&c)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
}
