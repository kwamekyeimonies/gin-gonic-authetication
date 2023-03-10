package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/gin-gonic-authetication/controller"
	"github.com/kwamekyeimonies/gin-gonic-authetication/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
}