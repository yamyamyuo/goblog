package main

import "github.com/gin-gonic/gin"

func getAbout(c *gin.Context) {
	c.String(200, "Hello")
}
