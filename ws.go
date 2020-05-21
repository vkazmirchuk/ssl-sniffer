package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func InitWebServer() {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/ws", serveWS)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

var upgrader = websocket.Upgrader{}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	WSConnections.Add(r.RemoteAddr, conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			WSConnections.Remove(r.RemoteAddr)
			break
		}
	}

}
