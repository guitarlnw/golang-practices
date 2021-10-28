package controller_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/guitarlnw/golang-practices/controller"
	mocks "github.com/guitarlnw/golang-practices/mocks/model"
	"github.com/guitarlnw/golang-practices/model"
	"github.com/stretchr/testify/suite"
)

type GetUserTestSuite struct {
	suite.Suite
	expectedData string
	app          *fiber.App
	mockModel    *mocks.IUserModel
}

func (suite *GetUserTestSuite) SetupTest() {
	suite.expectedData = `{"id":"","name":""}`
	suite.app = fiber.New()
	suite.mockModel = new(mocks.IUserModel)
}

func (suite *GetUserTestSuite) TestSuccess() {
	suite.mockModel.On("GetUserByID", "1", &model.UserModel{}).Return(nil)

	ctl := controller.UserCtl{
		UserModel: suite.mockModel,
	}

	suite.app.Get("/users/:id", ctl.GetUser)
	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, err := suite.app.Test(req)
	if err != nil {
		fmt.Println("Somthing went wrong!!")
	}
	body, _ := ioutil.ReadAll(resp.Body)

	suite.Equal(200, resp.StatusCode)
	suite.Equal(suite.expectedData, string(body))
}

func (suite *GetUserTestSuite) TestContentNotFound() {
	suite.expectedData = "content not found"
	suite.mockModel.On("GetUserByID", "1", &model.UserModel{}).Return(errors.New("content not found"))

	ctl := controller.UserCtl{
		UserModel: suite.mockModel,
	}

	suite.app.Get("/users/:id", ctl.GetUser)
	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, err := suite.app.Test(req)
	if err != nil {
		fmt.Println("Somthing went wrong!!")
	}
	body, _ := ioutil.ReadAll(resp.Body)

	suite.Equal(404, resp.StatusCode)
	suite.Equal(suite.expectedData, string(body))
}

func TestGetUserTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserTestSuite))
}
