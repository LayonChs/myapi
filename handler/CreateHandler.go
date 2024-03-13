package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}
