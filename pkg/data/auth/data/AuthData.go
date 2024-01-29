package data

import "go-crud-api/pkg/data/users/data"

type AuthData struct {
	User  data.UsersData `json:"user"`
	Token string         `json:"token"`
}
