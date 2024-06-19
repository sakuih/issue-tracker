package routes

import (
	"github.com/gin-gonic/gin"
	controller "issue-tracker.com/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	//incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/issue", controller.GetAllIssues())
	incomingRoutes.GET("/issue/:id", controller.GetIssue())
	incomingRoutes.POST("/issue", controller.PostIssue())
	incomingRoutes.DELETE("/issue/:id", controller.DeleteIssue())
	incomingRoutes.PUT("/issue/:id", controller.UpdateIssue())
}
