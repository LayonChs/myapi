package router

import (
	"myapi/handler"
	"github.com/gin-gonic/gin"
)

func initializeRouter(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowHandler)
		v1.POST("/opening", handler.CreateHandler)
		v1.DELETE("/opening", handler.DeleteHandler)
		v1.PUT("/opening", handler.UpdateHandler)
		v1.GET("/openings", handler.ListHandler)
	}

}



