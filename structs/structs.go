package structs

import (
	"github.com/segmentio/ksuid"
	"time"
)

type User struct {
	Id           ksuid.KSUID `json:"id"`
	Username     string      `json:"username"`
	Email        string      `json:"email"`
	PasswordHash string
	Created      time.Time `json:"created"`
	Permission   int8      `json:"permission"`
}
