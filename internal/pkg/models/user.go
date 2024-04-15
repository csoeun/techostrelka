package models

import "errors"

type user struct {
    Username string `json:"username"`
    Password string `json:"-"`
}

var userList = []user{
	
}

func registerNewUser(username, password string) (*user, error) {
    return nil, errors.New("placeholder error")
}

func isUsernameAvailable(username string) bool {
    return false
}