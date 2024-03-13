package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func ShowHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"msg":"GET Opening",
	})
}

func CreateHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"msg":"POST Opening",
	})
}

func DeleteHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"msg":"DELETE Opening",
	})
}

func UpdateHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"msg":"PUT Opening",
	})
}

func ListHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"msg":"GET Opening",
	})
}