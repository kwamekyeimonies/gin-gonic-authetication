package main

import "github.com/gin-gonic/gin"

func main(){

	server := gin.Default()

	server.GET(
		"/test",
		func(ctx *gin.Context) {
			ctx.JSON(200,gin.H{"Message":"Server running",})
		},
		)

	server.GET(
		"/retest",
		func(ctx *gin.Context) {
			ctx.JSON(200,gin.H{"Message":"Api Working,checking for live mode"})
		},
	)


	server.Run(":6666")
}