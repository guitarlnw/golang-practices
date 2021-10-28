package model_test

import (
	"errors"
	"testing"

	"github.com/guitarlnw/golang-practices/model"
	"github.com/stretchr/testify/suite"
)

type GetUserByIDModelTestSuite struct {
	suite.Suite
	expectedData model.UserModel
}

func (suite *GetUserByIDModelTestSuite) SetupTest() {
	suite.expectedData = model.UserModel{
		ID:   "1",
		Name: "Karn",
	}
}

func (suite *GetUserByIDModelTestSuite) TestSuccess() {
	var result model.UserModel
	m := model.NewUser()
	err := m.GetUserByID("1", &result)
	suite.Equal(nil, err)
	suite.Equal(suite.expectedData, result)
}

func (suite *GetUserByIDModelTestSuite) TestError() {
	var result model.UserModel
	m := model.NewUser()
	err := m.GetUserByID("2", &result)
	suite.Equal(errors.New("not found"), err)
	suite.Equal(model.UserModel{}, result)
}

func TestGetUserByIDModelTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserByIDModelTestSuite))
}
