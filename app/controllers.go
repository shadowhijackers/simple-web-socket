package app

import (
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Id string
	}{
		Id: "",
	}
	// id := uuid.New()

	data.Id = generateUId()
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, data)
}

func GangLocationHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("locations.html"))
	t.Execute(w, nil)
}
