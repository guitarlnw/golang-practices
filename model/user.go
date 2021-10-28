package model

import "errors"

type IUserModel interface {
	GetUserByID(id string, result *UserModel) error
}

type UserModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUser() IUserModel {
	return &UserModel{}
}

func (u UserModel) GetUserByID(id string, result *UserModel) error {
	// Should implement database logic here

	if id == "1" {
		result.ID = "1"
		result.Name = "Karn"
		return nil
	}
	return errors.New("not found")
}
