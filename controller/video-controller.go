package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/gin-gonic-framework/entity"
	"github.com/kwamekyeimonies/gin-gonic-framework/service"
)

type VideoController interface{
	FindAll() []entity.Video
	Save(ctx *gin.Context)
}

type controller struct{
	service service.Video_Service
}

func New(service service.Video_Service) VideoController{
	return controller{
		service: service,
	}
}