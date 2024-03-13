package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	initializeRouter(router)


	
	router.Run("localhost:8080")
}
