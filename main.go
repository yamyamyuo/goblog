package main

import "github.com/gin-gonic/gin"

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	archives := router.Group("/archives")
	{
		archives.GET("", allEntries)
		archives.GET(":id", entry)
		archives.POST(":id", newEntry)
		archives.DELETE(":id", deleteEntry)
	}
	about := router.Group("/about")
	{
		about.GET("", aboutPage)
	}

	router.Run("localhost:8000")
}
