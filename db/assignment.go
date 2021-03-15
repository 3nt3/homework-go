package db

import (
	"github.com/3nt3/homework/structs"
	"github.com/segmentio/ksuid"
)

func CreateAssignment(assignment structs.Assignment) (structs.Assignment, error) {
	id := ksuid.New()
	_, err := database.Exec("INSERT INTO assignments (id, content, course_id, due_date, creator_id, created_at, from_moodle) VALUES ($1, $2, $3, $4, $5, $6, $7)", id.String(), assignment.Title, assignment.Course, assignment.DueDate, assignment.Creator, assignment.Created, assignment.FromMoodle)

	newAssignment := assignment
	newAssignment.UID = id

	return newAssignment, err
}

func GetAssignmentByID(id string) (structs.Assignment, error) {
	row := database.QueryRow("SELECT * FROM assignments WHERE id = $1", id)

	var a structs.Assignment
	if row.Err() != nil {
		return a, row.Err()
	}

	err := row.Scan(&a.UID, &a.Title, &a.Course, &a.DueDate, &a.Creator, &a.Created, &a.FromMoodle)
	return a, err
}

func DeleteAssignment(id string) error {
	_, err := database.Exec("DELETE FROM assignments WHERE id = $1", id)
	return err
}