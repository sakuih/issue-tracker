package main

import (
  //"fmt"
  //"net/http"
  "github.com/gin-gonic/gin"
  helpers "issue-tracker.com/helpers"
  routes "issue-tracker.com/routes"
  "github.com/joho/godotenv"
  "log"
)

func main() {

  err := godotenv.Load(".env")
  if err != nil {
    log.Fatal(err)
  }
  helpers.ConnectToDb()
  gin.SetMode(gin.ReleaseMode)

  router := gin.New()
  router.Use(gin.Logger())
  //api := router.Group("/api")
  routes.UserRoutes(router)

  
  router.Run(":8000")


}



