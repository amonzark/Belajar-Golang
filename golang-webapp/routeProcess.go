package main

import (
	"html/template"
	"net/http"
)

// https://dasarpemrogramangolang.novalagung.com/B-form-value.html
func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("views/view.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("views/view.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var name = r.FormValue("name")
		var nick = r.FormValue("nick")
		var message = r.Form.Get("message")

		var data = map[string]string{"name": name, "nick": nick, "message": message}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
