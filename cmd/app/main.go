package main

import (
	"html/template"
	"log"
	"net/http"
)

type UserDetails struct { 
	Login string
	Password string
	Success bool
	StorageAccess string
}

var (
	tmpl = template.Must(template.ParseFiles("index.html"))
)

func handler(w http.ResponseWriter, req *http.Request) {
	data := UserDetails {
		Login: req.FormValue("login"),
		Password: req.FormValue("password"),
	}
	data.Success = true // если юзер уже есть в базе иначе false
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}