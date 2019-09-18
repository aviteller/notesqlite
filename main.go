package main

import (
	"fmt"
	"log"
	"net/http"

	"./controllers"
	m "./models"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"github.com/skratchdot/open-golang/open"
)

func main() {

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},                                               // All origins
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"}, // Allowing only get, just an example
	})

	database := m.GetDB()
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY autoincrement, title VARCHAR(255), message TEXT)")
	statement.Exec()

	router.HandleFunc("/api/notes", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/api/notes", controllers.GetNotes).Methods("GET")
	router.HandleFunc("/api/notes/{id}", controllers.DeleteNote).Methods("DELETE")
	router.HandleFunc("/api/notes/{id}", controllers.UpdateNote).Methods("PUT")

	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static/public").HTTPBox()))
	//router.PathPrefix("/workouts").Handler(http.FileServer(rice.MustFindBox("static/public").HTTPBox()))

	port := "9000"
	if port == "" {
		port = "8000"
	}

	err2 := open.Run("http://localhost:9000")
	if err2 != nil {
		log.Println(err2)
	}

	err := http.ListenAndServe(":"+port, c.Handler(router))
	if err != nil {
		fmt.Println(err)
	}

}
