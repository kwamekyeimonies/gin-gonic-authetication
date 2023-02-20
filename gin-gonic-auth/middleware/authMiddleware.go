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
	}
}