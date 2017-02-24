package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/asdine/storm"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	s := &server{}

	rtr := httprouter.New()
	rtr.HandlerFunc("POST", "/register", s.Register)
	rtr.ServeFiles("/public/*filepath", http.Dir("public"))
	db, err := storm.Open("watchr.db")
	if err != nil {
		log.Fatalln(err)
	}

	s.db = db
	s.router = rtr
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"X-API"},
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
	}).Handler(s.router)

	// http.Handle("/app", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", handler)
}

// Register takes the HTTP request and attempts to create a user
func (s *server) Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")
	user := &User{Name: name, Email: email}

	err := s.db.Save(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("User registration successful")
	t, err := template.ParseFiles("public/register.html")
	if err != nil {
		panic(err)
	}
	log.Println(user)
	t.Execute(w, user)
}

type server struct {
	router *httprouter.Router
	db     *storm.DB
}

// User is an object containing an email and a name.
type User struct {
	ID    int `storm:"id,increment"`
	Name  string
	Email string
}
