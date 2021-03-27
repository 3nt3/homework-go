package structs

import (
	"github.com/segmentio/ksuid"
	"time"
)

// max age in days
const MaxSessionAge int = 90

type User struct {
	ID           ksuid.KSUID `json:"id"`
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	PasswordHash string
	Created      time.Time `json:"created"`
	Permission   int8      `json:"permission"`
	Courses      []Course  `json:"courses"`
	MoodleURL    string    `json:"moodle_url"`
	MoodleToken  string    `json:"moodle_token"`
	MoodleUserID int       `json:"moodle_user_id"`
}

type CleanUser struct {
	ID           ksuid.KSUID `json:"id"`
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	Created      time.Time   `json:"created"`
	Permission   int8        `json:"permission"`
	Courses      []Course    `json:"courses"`
	MoodleURL    string      `json:"moodle_url"`
	MoodleUserID int         `json:"moodle_user_id"`
}

func (u User) GetClean() CleanUser {
	return CleanUser{
		ID:           u.ID,
		Username:     u.Username,
		Email:        u.Email,
		Created:      u.Created,
		Permission:   u.Permission,
		Courses:      u.Courses,
		MoodleURL:    u.MoodleURL,
		MoodleUserID: u.MoodleUserID,
	}
}

func (a Assignment) GetClean() CleanAssignment {
	return CleanAssignment{
		UID:       	a.UID,
		Creator:    a.Creator.GetClean(),
		Created:    a.Created,
		Title:      a.Title,
		DueDate:    a.DueDate,
		Course:     a.Course,
		FromMoodle: a.FromMoodle,
	}
}

type Session struct {
	UID     ksuid.KSUID `json:"uid"`
	UserID  ksuid.KSUID `json:"user_id"`
	Created time.Time   `json:"created"`
}

type Assignment struct {
	UID        ksuid.KSUID `json:"uid"`
	Creator    User        `json:"creator"`
	Created    time.Time   `json:"created"`
	Title      string      `json:"title"`
	DueDate    time.Time   `json:"due_date"`
	Course     int         `json:"course"`
	FromMoodle bool        `json:"from_moodle"`
}

type CleanAssignment struct {
	UID        ksuid.KSUID `json:"uid"`
	Creator    CleanUser        `json:"creator"`
	Created    time.Time   `json:"created"`
	Title      string      `json:"title"`
	DueDate    time.Time   `json:"due_date"`
	Course     int         `json:"course"`
	FromMoodle bool        `json:"from_moodle"`
}

type Course struct {
	ID          interface{}  `json:"id"`
	Name        string       `json:"name"`
	Teacher     string       `json:"teacher"`
	FromMoodle  bool         `json:"from_moodle"`
	Assignments []Assignment `json:"asssignments"`
	User        ksuid.KSUID  `json:"user"`
}

type CleanCourse struct {
	ID          interface{}  `json:"id"`
	Name        string       `json:"name"`
	Teacher     string       `json:"teacher"`
	FromMoodle  bool         `json:"from_moodle"`
	Assignments []CleanAssignment `json:"asssignments"`
	User        ksuid.KSUID  `json:"user"`
}

func (c Course) GetClean() CleanCourse {
	cc := CleanCourse{
		ID:          c.ID,
		Name:        c.Name,
		Teacher:     c.Teacher,
		FromMoodle:  c.FromMoodle,
		User:        c.User,
	}
	for i := 0; i < len(c.Assignments); i++ {
		cc.Assignments = append(cc.Assignments, c.Assignments[i].GetClean())
	}

	return cc
}

type CachedCourse struct {
	ID ksuid.KSUID `json:"id"`
	Course
	MoodleURL string      `json:"moodle_url"`
	UserID    ksuid.KSUID `json:"user_id"`
	CachedAt  time.Time   `json:"cached_at"`
}
