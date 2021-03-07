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

	err := db.InitDatabase()
	if err != nil {
		logging.ErrorLogger.Printf("error connecting to db: %v\n", err)
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/user", routes.User)

	logging.ErrorLogger.Fatalln(http.ListenAndServe(":8000", r).Error())
}
