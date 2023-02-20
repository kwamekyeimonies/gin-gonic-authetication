package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/gin-gonic-authetication/utils"
)


func Authenticate() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("token")
		if clientToken == ""{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":fmt.Sprintf("No Authorization Header provided")})
			ctx.Abort()
			return
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != ""{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err})
			ctx.Abort()
			return
		}

		ctx.Set("email",claims.Email)
		ctx.Set("first_name",claims.Email)
		ctx.Set("last_name",claims.Last_Name)
		ctx.Set("uid",claims.Uid)
		ctx.Set("user_type",claims.User_Type)
		ctx.Next()
	}
}