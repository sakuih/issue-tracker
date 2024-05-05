package routes

import (
  controller "issue-tracker.com/controllers"
  "github.com/gin-gonic/gin"
)

func UserRoutes (incomingRoutes *gin.Engine) {
  //incomingRoutes.Use(middleware.Authenticate())
  incomingRoutes.GET("/issues", controller.GetIssues())
}


