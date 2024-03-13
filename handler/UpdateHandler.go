package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func UpdateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}
