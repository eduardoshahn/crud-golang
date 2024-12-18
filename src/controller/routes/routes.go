package routes

import (
	"github.com/eduardoshahn/crud-golang.git/src/controller"
	"github.com/gin-gonic/gin"
)

// Func que irá inicializar as rotas
func InitRoutes(r *gin.RouterGroup){
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	sr.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}