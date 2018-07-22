package controller

import ("github.com/gin-gonic/gin")

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": "rakki",
		"cant_post": 4000,
	})
}