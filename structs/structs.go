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
}

type CleanUser struct {
	ID         ksuid.KSUID `json:"id"`
	Username   string      `json:"username"`
	Email      string      `json:"email"`
	Created    time.Time   `json:"created"`
	Permission int8        `json:"permission"`
}

func (u User) GetClean() CleanUser {
	return CleanUser{
		ID:         u.ID,
		Username:   u.Username,
		Email:      u.Email,
		Created:    u.Created,
		Permission: u.Permission,
	}
}

type Session struct {
	UID     ksuid.KSUID `json:"uid"`
	UserID  ksuid.KSUID `json:"user_id"`
	Created time.Time   `json:"created"`
}
