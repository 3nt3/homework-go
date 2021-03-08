package routes

import (
	"encoding/json"
	"github.com/3nt3/homework/db"
	"github.com/3nt3/homework/logging"
	"github.com/gorilla/mux"
	"net/http"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	var userData map[string]string
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		logging.WarningLogger.Printf("error decoding request: %v\n", err)
		_ = returnApiResponse(w, apiResponse{Content: nil, Errors: []string{"invalid request"}}, 400)
		return
	}

	username, ok := userData["username"]
	if !ok {
		logging.WarningLogger.Printf("error decoding request: field 'username' does not exist\n")
		_ = returnApiResponse(w, apiResponse{Content: nil, Errors: []string{"invalid request"}}, 400)
		return
	}

	email, ok := userData["email"]
	if !ok {
		logging.WarningLogger.Printf("error decoding request: field 'email' does not exist\n")
		_ = returnApiResponse(w, apiResponse{Content: nil, Errors: []string{"invalid request"}}, 400)
		return
	}

	password, ok := userData["username"]
	if !ok {
		logging.WarningLogger.Printf("error decoding request: field 'password' does not exist\n")
		_ = returnApiResponse(w, apiResponse{Content: nil, Errors: []string{"invalid request"}}, 400)
		return
	}

	user, err := db.NewUser(username, email, password)
	if err != nil {
		logging.ErrorLogger.Printf("error creating new user: %v\n", err)
		_ = returnApiResponse(w, apiResponse{Content: nil, Errors: []string{"internal server error"}}, 500)
		return
	}

	_ = returnApiResponse(w, apiResponse{Content: user.GetClean(), Errors: []string{}}, 200)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		logging.WarningLogger.Printf("no id specified\n")
		_ = returnApiResponse(w, apiResponse{Content: nil, Errors: []string{"invalid request"}}, 400)
		return
	}

	user, err := db.GetUserById(id)
	if err != nil {
		logging.ErrorLogger.Printf("error fetching user from db: %v\n", err)
		_ = returnApiResponse(w, apiResponse{Content: nil, Errors: []string{"internal server error"}}, 500)
		return
	}

	_ = returnApiResponse(w, apiResponse{Content: user, Errors: []string{}}, 200)
}
