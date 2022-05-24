package main

import (
	// "fmt"
	// "html"
	// "log"
	// "net/http"
	"github.com/gin-gonic/gin"
)

// CURD
func GetAllUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ping !",
	})
}

func main() {

	router := gin.Default()

	router.GET("/", GetAllUser)

	router.Run(":8081")

}
