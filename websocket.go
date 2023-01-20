package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var players = make(map[*websocket.Conn]bool)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Add the new player to the map of players
	players[ws] = true

	for {
		// Handle messages sent by the player
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(players, ws)
			break
		}

		updatePlayers()
	}
}

func updatePlayers() {
	for player := range players {
		err := player.WriteJSON(players)
		if err != nil {
			log.Println(err)
			player.Close()
			delete(players, player)
		}
	}
}
