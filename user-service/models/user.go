package models

type User struct {
	Id       int    `json:"Id"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
