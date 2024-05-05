package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func main() {

  api := router.Group("/api")

  router := gin.New()
  router.Use(gin.Logger())


  routes.UserRoutes(router)

  api.GET("/", getData)


}

func getData(c *gin.Context) {
  c.Header("Content-Type", "application/json")

  c.JSON(http.StatusOK, gin.H {
    "message": "GetData"
  })


}





