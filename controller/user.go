package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guitarlnw/golang-practices/model"
)

type IUser interface {
	GetUser(c *fiber.Ctx) error
}

type User struct {
	UserModel model.IUser
}

func NewUser() IUser {
	return &User{
		UserModel: model.NewUser(),
	}
}

func (u *User) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON("Bad request")
	}
	user := model.User{}
	if err := u.UserModel.GetUserByID(id, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(user)
}
