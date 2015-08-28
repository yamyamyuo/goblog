package main

import "github.com/gin-gonic/gin"

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// archive := router.Group("/archive")
	// {
	// 	archive.GET("", getAllEntries)
	// 	archive.GET("id", getEntry)
	// 	archive.POST("id", postEntry)
	// }
	about := router.Group("/about")
	{
		about.GET("", getAbout)
	}

	router.Run("localhost:8000")
}
