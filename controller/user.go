package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guitarlnw/golang-practices/model"
)

type IUserCtl interface {
	GetUser(c *fiber.Ctx) error
}

type UserCtl struct {
	UserModel model.IUserModel
}

func NewUser() IUserCtl {
	return &UserCtl{
		UserModel: model.NewUser(),
	}
}

func (u *UserCtl) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON("Bad request")
	}
	user := model.UserModel{}
	if err := u.UserModel.GetUserByID(id, &user); err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.Status(200).JSON(user)
}
