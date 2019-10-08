package main

import (
	"fmt"
	"net/http"

	"./controllers"
	m "./models"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},                                               // All origins
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"}, // Allowing only get, just an example
	})

	m.InitTables()

	router.HandleFunc("/api/notes", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/api/notes", controllers.GetNotes).Methods("GET")
	router.HandleFunc("/api/notes/{id}", controllers.DeleteNote).Methods("DELETE")
	router.HandleFunc("/api/notes/{id}", controllers.UpdateNote).Methods("PUT")

	router.HandleFunc("/api/workouts", controllers.CreateWorkout).Methods("POST")
	router.HandleFunc("/api/workouts", controllers.GetWorkouts).Methods("GET")
	router.HandleFunc("/api/workouts/{id}", controllers.GetWorkoutDetails).Methods("GET")
	router.HandleFunc("/api/workoutexport/{id}", controllers.ExportWorkout).Methods("GET")
	router.HandleFunc("/api/workouts/{id}", controllers.DeleteWorkout).Methods("DELETE")
	router.HandleFunc("/api/workouts/{id}", controllers.UpdateWorkout).Methods("PUT")

	router.HandleFunc("/api/actions", controllers.CreateAction).Methods("POST")
	router.HandleFunc("/api/workouts/actions/{workoutid}", controllers.GetActionsForWorkout).Methods("GET")
	router.HandleFunc("/api/actions/swap/{firstID}/{secondID}", controllers.SwapActionPos).Methods("GET")
	router.HandleFunc("/api/actions/{id}", controllers.DeleteAction).Methods("DELETE")
	router.HandleFunc("/api/actions/{id}", controllers.UpdateAction).Methods("PUT")

	router.HandleFunc("/api/budgets", controllers.CreateBudget).Methods("POST")
	router.HandleFunc("/api/budgets/{month}", controllers.GetBudgets).Methods("GET")
	router.HandleFunc("/api/budgets/stats/{month}", controllers.GetBudgetStatsByMonth).Methods("GET")
	router.HandleFunc("/api/budgets/{id}", controllers.DeleteBudget).Methods("DELETE")
	router.HandleFunc("/api/budgets/{id}", controllers.UpdateBudget).Methods("PUT")

	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("static/public").HTTPBox()))
	//router.PathPrefix("/workouts").Handler(http.FileServer(rice.MustFindBox("static/public").HTTPBox()))

	port := "9000"
	if port == "" {
		port = "8000"
	}

	// err2 := open.Run("http://localhost:9000")
	// if err2 != nil {
	// 	log.Println(err2)
	// }

	err := http.ListenAndServe(":"+port, c.Handler(router))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("server running")
	}

}
