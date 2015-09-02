package main

import "github.com/gin-gonic/gin"

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLFiles("index.html")
	router.StaticFile("/simple-sidebar.css", "./ui/css/simple-sidebar.css")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	archives := router.Group("/archives")
	{
		archives.GET("", allEntries)
		archives.GET(":id", entry)
		archives.POST("", newEntry)
		archives.DELETE(":id", deleteEntry)
	}
	about := router.Group("/about")
	{
		about.GET("", aboutPage)
	}

	router.Run("localhost:8000")
}
