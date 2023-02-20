package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/gin-gonic-authetication/models"
	"github.com/kwamekyeimonies/gin-gonic-authetication/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if err := utils.CheckUserType(ctx, "ADMIN"); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"eror": err.Error()})
			return
		}

		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(ctx.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}

		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(ctx.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{
			{"_id", bson.D{{"_id", "null"}}},
			{"total_count", bson.D{{"$sum", 1}}},
			{"data", bson.D{{"$push", "$$ROOT"}}}}}}
		projectStage := bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}}}}}
		result, err := UserCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		}
		var allUsers []bson.D
		if err = result.All(c, &allUsers); err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, allUsers[0])
	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userId := ctx.Param("user_id")
		if err := utils.MatchUserTypeToUid(ctx, userId); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := UserCollection.FindOne(c, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)

	}
}
