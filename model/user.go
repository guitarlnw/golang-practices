package model

import "errors"

type IUser interface {
	GetUserByID(id string, result *User) error
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUser() IUser {
	return &User{}
}

func (u User) GetUserByID(id string, result *User) error {
	// Should implement database logic here

	if id == "1" {
		result.ID = "1"
		result.Name = "2"
		return nil
	}
	return errors.New("not found")
}
