package main

import (
	"auth/handle"
	"auth/model"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	_ "github.com/lib/pq"
	"path/filepath" // so that we can make path joins compatible on all OS
)

var tmpl = template.New("")

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	// tmpl:= template.Must(template.ParseFiles("templates/index.html"))
	// tmpl.Execute(w, nil)
	//tmpl.ExecuteTemplate(w, "fstyle.css",nil)
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func main() {
	
  	
	_, err := tmpl.ParseGlob(filepath.Join(".", "templates", "*.html"))
	if err != nil {
		log.Fatalf("Unable to parse templates: %v\n", err)
	}
	// _, err = tmpl.ParseGlob(filepath.Join(".", "templates/css", "*.css"))
	// if err != nil {
	// 	log.Fatalf("Unable to parse templates: %v\n", err)
	// }

	fmt.Println(filepath.Join(".", "templates", "*.html"))
	fmt.Println(filepath.Join(".", "templates", "*.css"))

	fs:= http.FileServer(http.Dir("templates/"))
	// Registering routes and handler that we will implement
	//multi := http.NewServeMux() 
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/signin", handle.Signin).Methods("POST")
	r.HandleFunc("/signup", handle.Signup).Methods("POST")
	r.HandleFunc("/", handler).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	
	//multi.Handle("/static/", http.StripPrefix("/static", fs))
	// multi.HandleFunc("/", handler)
	// initialize our database connection


	db := model.InitDB()
	defer db.Close()
	// start the server on port 8000
	fmt.Println("Listening and serving.....")
	log.Fatal(http.ListenAndServe(":8000", r))
}

