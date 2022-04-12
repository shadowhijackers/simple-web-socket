package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type AbstractApp interface {
	Initialize()
	SetupRouter()
	ListenSocket()
}

type App struct {
	Router   *mux.Router
	Upgrader *websocket.Upgrader
	Conns    *GangConnections
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.Upgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	app.Conns = &GangConnections{}
	app.Conns.gangs = make(map[GangId]GangConnection)
	app.SetupRouter()
	app.ListenSocket()
}

func (app *App) SetupRouter() {
	app.Get("/", IndexHandler)
	app.Get("/gang/locations/{id}", GangLocationHandler)
}

func (app *App) Run(host string) {
	err := http.ListenAndServe(host, app.Router)
	handleError(err)
	log.Println("Application Started Running")
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("Get")
}

func (app *App) Post(path string, f http.HandlerFunc) {
	app.Router.HandleFunc(path, f).Methods("POST")
}
