package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type GangId = string
type IP = string
type GangConnection map[IP]*websocket.Conn

type GangConnections struct {
	gangs map[GangId]GangConnection
}

func (gangConnections *GangConnections) register(conn *websocket.Conn, r *http.Request, msg string, gangConns chan GangConnection) {
	param := mux.Vars(r)
	gangId := param["id"]
	ip := conn.RemoteAddr().String()

	if gangConnections.gangs[gangId] != nil {
		// gangConnections.gangs[gangId] = make(GangConnection)
		gangConnections.gangs[gangId][ip] = conn
	} else {
		gangConnections.gangs[gangId] = make(GangConnection)
		gangConnections.gangs[gangId][ip] = conn
	}

	gangConns <- gangConnections.gangs[gangId]
	close(gangConns)
}

func broadCast(conn *websocket.Conn, msgType int, msg []byte, gangConns GangConnection) {
	for ip, bcconn := range gangConns {
		if ip == conn.RemoteAddr().String() {
			if err := bcconn.WriteMessage(msgType, []byte("Testing... Sent message from You "+conn.RemoteAddr().String()+" "+string(msg))); err != nil {
				return
			}
			continue
		}
		// Write message back to browser
		if err := bcconn.WriteMessage(msgType, []byte("Testing... Recived message from "+conn.RemoteAddr().String()+" "+string(msg))); err != nil {
			return
		}
	}
}

func (app *App) ListenSocket() {
	app.Router.HandleFunc("/gang/{id}", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := app.Upgrader.Upgrade(w, r, nil)
		for {
			msgType, msg, err := conn.ReadMessage()

			if err != nil {
				return

			}

			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			gangConns := make(chan GangConnection)
			conns := make(GangConnection)
			go app.Conns.register(conn, r, string(msg), gangConns)
			conns = <-gangConns
			go broadCast(conn, msgType, msg, conns)
		}
	})
}
