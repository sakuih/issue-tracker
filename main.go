package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()
	r := gin.Default()

	api := r.Group("/api")

	api.GET("/issues", GetIssues)
	api.POST("/issues", CreateIssue)
	api.GET("/issues/:id", GetIssue)
	api.PUT("/issues/:id", UpdateIssue)
	api.DELETE("/issues/:id", DeleteIssue)

	r.Run(":8080")

}
