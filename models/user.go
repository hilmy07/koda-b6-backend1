package models

type Users struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var ListUser []Users
var UserCounter int