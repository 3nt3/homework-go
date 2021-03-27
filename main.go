package main

import (
	"github.com/3nt3/homework/db"
	"github.com/3nt3/homework/logging"
	"github.com/3nt3/homework/routes"
	"github.com/gorilla/mux"
	"net/http"
)


func main() {
	logging.InitLoggers()

	err := db.InitDatabase(false)
	if err != nil {
		logging.ErrorLogger.Printf("error connecting to db: %v\n", err)
		return
	}

	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(routes.HandleCORSPreflight)
	r.HandleFunc("/user/register", routes.NewUser).Methods("POST")
	r.HandleFunc("/user", routes.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserById).Methods("GET")
	r.HandleFunc("/user/login", routes.Login).Methods("POST")
	r.HandleFunc("/assignment", routes.CreateAssignment).Methods("POST")
	r.HandleFunc("/assignment", routes.DeleteAssignment).Methods("DELETE")
	r.HandleFunc("/course/active", routes.GetActiveCourses)
	r.HandleFunc("/moodle/authenticate", routes.MoodleAuthenticate).Methods("POST")

	logging.InfoLogger.Println("started server on port :8000")
	logging.ErrorLogger.Fatalln(http.ListenAndServe(":8000", r).Error())
}

