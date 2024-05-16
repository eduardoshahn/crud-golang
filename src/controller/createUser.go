package controller

import (
	"net/http"

	"github.com/eduardoshahn/crud-golang.git/src/configuration/logger"
	"github.com/eduardoshahn/crud-golang.git/src/configuration/validation"
	"github.com/eduardoshahn/crud-golang.git/src/controller/model/request"
	"github.com/eduardoshahn/crud-golang.git/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context){
	logger.Info("Init createUser controller",
		zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
		zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.UserResponse{
		ID: "test",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	logger.Info("user create successfully",
	zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
