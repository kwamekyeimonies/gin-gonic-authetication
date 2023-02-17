package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/gin-gonic-authetication/models"
	"github.com/kwamekyeimonies/gin-gonic-authetication/utils"
	"go.mongodb.org/mongo-driver/bson"
)


func GetUsers(){

}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId := ctx.Param("user_id")
		if err := utils.MatchUserTypeToUid(ctx, userId); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var c,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := UserCollection.FindOne(c,bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err != nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)

	}
}