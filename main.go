package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/register", Register)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

// Register takes the HTTP request and attempts to create a user
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

// User is an object containing an email and a name.
type User struct {
	Name  string
	Email string
}
