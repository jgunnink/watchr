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
	value := r.FormValue("user")
	user := &User{Name: value}
	t, err := template.ParseFiles("public/register.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, user)
}

type User struct {
	Name string
}
