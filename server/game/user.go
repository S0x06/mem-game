package game

import "errors"

type User struct {
	ID string `json:"id"`
}

type UserFact struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

var ErrCantFindUser = errors.New("can't find user")
