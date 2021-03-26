package db

import (
	"encoding/json"
	"github.com/3nt3/homework/structs"
)

func GetUserCachedCourses(user structs.User) ([]structs.CachedCourse, error) {
	var courses []structs.CachedCourse

	// get all cached moodle courses where moodle_url == the users moodle url and the userID == user.id
	query := "SELECT * FROM moodle_cache WHERE moodle_url == $1 AND user_id == $2"
	rows, err := database.Query(query, user.MoodleURL, user.ID)
	if err != nil {
		return nil, err
	}

	// iterate through rows
	for rows.Next() {
		// create variables
		var newCourse structs.CachedCourse
		var jsonString string

		// populate variables
		err = rows.Scan(&newCourse.ID, &jsonString, &newCourse.MoodleURL, &newCourse.CachedAt, &newCourse.UserID)
		if err != nil {
			return nil, err
		}

		// decode json encoded course data
		if err = json.Unmarshal([]byte(jsonString), &newCourse.Course); err != nil {
			return nil, err
		}

		// append to array
		courses = append(courses, newCourse)
	}

	// return
	return courses, nil
}
