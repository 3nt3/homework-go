package routes

import (
	"bytes"
	"encoding/json"
	"github.com/3nt3/homework/db"
	"github.com/3nt3/homework/logging"
	"github.com/3nt3/homework/structs"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateAssignment(t *testing.T) {
	logging.InitLoggers()

	// initialize database
	err := db.InitDatabase(true)
	if err != nil {
		t.Errorf("error initializing database")
	}

	// register?
	body, _ := json.Marshal(map[string]string{
		"username": "test",
		"email": "test@example.com",
		"password": "test123",
	})

	req, err := http.NewRequest("POST", "http://localhost:8000/user", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("error requesting: %v", err)
	}
	rr := httptest.NewRecorder()

	NewUser(rr, req)

	result := rr.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("request failed with status code %d", result.StatusCode)
	}

	resp := apiResponse{}
	err = json.NewDecoder(result.Body).Decode(&resp)
	if err != nil {
		t.Errorf("error decoding body: %v", err)
	}

	if len(resp.Errors) > 0 {
		t.Errorf("request failed with errors: %v", resp.Errors)
	}

	logging.InfoLogger.Printf("user: %+v\n", resp.Content)

	a := structs.Assignment{
		Title:      "test assignment",
		DueDate:    time.Date(2021, 11, 16, 0, 0, 0, 0, time.UTC),
		Course:     123,
		FromMoodle: false,
	}
	body, _ = json.Marshal(a)

	req, err = http.NewRequest("POST", "http://localhost:8000/assignment", bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("error requesting: %v", err)
	}
	arr := httptest.NewRecorder()

	aResult := arr.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("request failed with status code %d", result.StatusCode)
	}

	aResp := apiResponse{}
	err = json.NewDecoder(aResult.Body).Decode(&aResp)
	if err != nil {
		t.Errorf("error decoding body: %v", err)
	}

	logging.InfoLogger.Printf("assignment")

	_ = db.DropTables()
}
