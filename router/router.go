package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{"Message":"Teste Okay"})
	})
	router.Run("localhost:8080")
}
