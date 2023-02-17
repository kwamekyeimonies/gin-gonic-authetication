package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kwamekyeimonies/gin-gonic-authetication/database"
	"github.com/kwamekyeimonies/gin-gonic-authetication/models"
	"github.com/kwamekyeimonies/gin-gonic-authetication/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var validate = validator.New()
var UserCollection *mongo.Collection = database.OpenCollection(database.Client,"user")

func Login() gin.HandlerFunc{
	return func(ctx *gin.Context) {

		var c,cancel = context.WithTimeout(context.Background(),100*time.Second)
		var user models.User
		var foundUser models.User

		if err := ctx.BindJSON(&user); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		err := UserCollection.FindOne(c,bson.M{"email":user.Email}).Decode(&foundUser)
		defer cancel()

		if err != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Email or Password Inccorrect"})
			return
		}

		passwordIsValid, msg := utils.VerifyPassword(*&user.Password, *&foundUser.Password)
		defer cancel()

		if passwordIsValid != true{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}

		if foundUser.Email == nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"User not found"})
			return
		}

		token,refreshToken,_ :=utils.GenerateAllTokens(foundUser.Email,foundUser.First_name,foundUser.Last_name,foundUser.User_type,foundUser.User_id)
		utils.U


	}

}


func SignUp() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		
		var c,cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		if err := ctx.BindJSON(&user); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return
		}
		count,err := UserCollection.CountDocuments(c,bson.M{"email":user.Email})
		defer cancel()
		if err != nil{
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		}

		password := HashPassword(user.Password)
		user.Password = &password

		count, err = UserCollection.CountDocuments(c,bson.M{"phone":user.Phone})
		defer cancel()
		if err != nil{
			log.Panic(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		if count > 0 {
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":"Email or Phone Number Already Exist"})
		}
		user.Created_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token,refreshToken,_ := utils.GenerateAllTokens(user.Email,user.First_name,user.Last_name,user.User_type,user.User_id )
		user.Token = token
		user.Refresh_Token = refreshToken

		resultInsertionNumber, insertErr := UserCollection.InsertOne(c,user)
		if insertErr != nil{
			msg := fmt.Sprintf("User Item was not created")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		ctx.JSON(http.StatusOK, resultInsertionNumber)
	}
}