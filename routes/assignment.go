package routes

import (
	"database/sql"
	"encoding/json"
	"github.com/3nt3/homework/db"
	"github.com/3nt3/homework/logging"
	"github.com/3nt3/homework/structs"
	"net/http"
)

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	user, authenticated, err := getUserBySession(r)

	if err != nil {
		logging.ErrorLogger.Printf("error getting user by session: %v\n", err)
		_ = returnApiResponse(w, apiResponse{
			Content: nil,
			Errors:  []string{"internal server error"},
		}, 500)
		return
	}

	if !authenticated {
		_ = returnApiResponse(w, apiResponse{
			Content: nil,
			Errors:  []string{"invalid session"},
		}, 401)
		return
	}

	var assignment structs.Assignment
	err = json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {
		logging.WarningLogger.Printf("error decoding: %v\n", err)
		_ = returnApiResponse(w, apiResponse{
			Content: nil,
			Errors: []string{"bad request"},
		}, http.StatusBadRequest)
		return
	}

	assignment.Creator = user.ID

	assignment, err = db.CreateAssignment(assignment)
	if err != nil {
		logging.ErrorLogger.Printf("error creating assignment: %v\n", err)
		_ = returnApiResponse(w, apiResponse{
			Content: nil,
			Errors: []string{"internal server error"},
		}, http.StatusInternalServerError)
		return
	}

	_ = returnApiResponse(w, apiResponse{
		Content: assignment,
		Errors: []string{},
	}, http.StatusOK)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	keyValue, ok := r.URL.Query()["id"]
	if !ok {
		_ = returnApiResponse(w, apiResponse{
			Content: nil,
			Errors: []string{"bad request"},
		}, http.StatusBadRequest)
		return
	}

	user, authenticated, err := getUserBySession(r)

	if err != nil {
		logging.ErrorLogger.Printf("error getting user by session: %v\n", err)
		_ = returnApiResponse(w, apiResponse{
			Content: nil,
			Errors:  []string{"internal server error"},
		}, 500)
		return
	}

	if !authenticated {
		_ = returnApiResponse(w, apiResponse{
			Content: nil,
			Errors:  []string{"invalid session"},
		}, 401)
		return
	}

	assignment, err := db.GetAssignmentByID(keyValue[1])
	if err != nil {
		if err == sql.ErrNoRows {
			_ = returnApiResponse(w, apiResponse{
				Content: nil,
				Errors:  []string{"not found"},
			}, http.StatusNotFound)
			return
		}

		_ = returnApiResponse(w, apiResponse{
			Content: nil,
			Errors:  []string{"internal server error"},
		}, http.StatusInternalServerError)
		return
	}

	if assignment.Creator != user.ID {
		_ = returnApiResponse(w, apiResponse{
			Content: nil,
			Errors:  []string{"you are not the creator of this assignment"},
		}, http.StatusForbidden)
		return
	}

	err = db.DeleteAssignment(assignment.UID.String())
	if err != nil {
		_ = returnApiResponse(w, apiResponse{
			Content: nil,
			Errors:  []string{"internal server error"},
		}, http.StatusInternalServerError)
		return
	}

	_ = returnApiResponse(w, apiResponse{
		Content: assignment,
		Errors: []string{},
	}, http.StatusOK)
}