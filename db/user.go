package db

import (
	"github.com/3nt3/homework/structs"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func NewUser(username string, email string, password string) (structs.User, error) {
	id := ksuid.New()

	hash, err := hashPassword(password)
	if err != nil {
		return structs.User{}, err
	}

	_, err = database.Exec("insert into users (id, username, email, password_hash, permission, created_at) values ($1, $2, $3, $4, $5, 0);", id.String(), username, email, hash)
	if err != nil {
		return structs.User{}, err
	}

	return structs.User{
		Id:           id,
		Username:     username,
		Email:        email,
		PasswordHash: hash,
		Created:      time.Now(),
		Permission:   0,
	}, nil
}

func GetUserByUsername(username string) (structs.User, error) {
	row := database.QueryRow("select * from users where username = $1;", username)
	if row.Err() != nil {
		return structs.User{}, row.Err()
	}

	var user structs.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.Created, &user.Permission)
	if err != nil {
		return structs.User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string ) (structs.User, error) {
	row := database.QueryRow("select * from users where email = $1;", email)
	if row.Err() != nil {
		return structs.User{}, row.Err()
	}

	var user structs.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.Created, &user.Permission)
	if err != nil {
		return structs.User{}, err
	}

	return user, nil
}

func GetUserById(id string ) (structs.User, error) {
	row := database.QueryRow("select * from users where id = $1;", id)
	if row.Err() != nil {
		return structs.User{}, row.Err()
	}

	var user structs.User
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.PasswordHash, &user.Created, &user.Permission)
	if err != nil {
		return structs.User{}, err
	}

	return user, nil
}

func Authenticate(username string, password string) (structs.User, bool, error) {
	// get user by username
	user, err := GetUserByUsername(username)

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		// if incorrect password, return
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return structs.User{}, false, nil
		}

		// if other error return error
		return structs.User{}, false, err
	}

	// if no error, return authenticated
	return structs.User{}, true, nil
}


func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
