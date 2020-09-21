package main

import "github.com/gin-gonic/gin"
import "net/http"

func main() {
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"message": "ok boy"})
	})
	r.Run()
}
