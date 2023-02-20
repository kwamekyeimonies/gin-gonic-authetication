package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/gin-gonic-authetication/controller"
)


func AuthRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.POST("/users/signup",controller.SignUp())
	incomingRoutes.POST("/users/login",controller.Login())

}