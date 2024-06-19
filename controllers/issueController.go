package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetAllIssues() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"html": "Success"})
	}

}

func GetIssue() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UpdateIssue() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func PostIssue() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteIssue() gin.HandlerFunc {
	return func(c *gin.Context) {

	}

}
