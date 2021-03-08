package routes

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Content interface{} `json:"content"`
	Errors []string `json:"errors"`
}

func returnApiResponse(w http.ResponseWriter, response apiResponse, status int) error {
	err := json.NewEncoder(w).Encode(response)
	w.WriteHeader(status)
	return err
}
