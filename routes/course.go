package routes

import (
	"github.com/3nt3/homework/db"
	"github.com/3nt3/homework/logging"
	"github.com/3nt3/homework/structs"
	"net/http"
)

// TODO: return all courses with assignments whose due dates lie in the future
func GetActiveCourses(w http.ResponseWriter, r *http.Request) {
	HandleCORSPreflight(w, r)

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

	courses, err := db.GetMoodleUserCourses(user)
	if err != nil {
		logging.ErrorLogger.Printf("error: %v\n", err)
	}

	var cleanCourses []structs.CleanCourse
	for _, c := range courses {
		cleanCourses = append(cleanCourses, c.GetClean())
	}

	_ = returnApiResponse(w, apiResponse{
		Content: cleanCourses,
		Errors:  []string{},
	}, 200)
}

// TODO
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	HandleCORSPreflight(w, r)

}
