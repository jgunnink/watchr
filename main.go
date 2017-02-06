package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/register", Register)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("user")
	email := r.FormValue("email")
	user := &User{Name: name, Email: email}
	t, err := template.ParseFiles("public/register.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, user)
}

type User struct {
	Name  string
	Email string
}
