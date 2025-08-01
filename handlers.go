package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Issue struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Status   string `json:"status"`
	Severity int    `json:"severity"`
}

func CreateIssue(c *gin.Context) {
	var issue Issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	res, err := DB.Exec("INSERT INTO issues (title, content, status, severity) VALUES (?, ?, ?, ?)",
		issue.Title, issue.Content, issue.Status, issue.Severity)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	id, _ := res.LastInsertId()
	issue.ID = int(id)
	c.JSON(201, issue)
}

func GetIssues(c *gin.Context) {
	status := c.Query("status")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	query := "SELECT id, title, content, status FROM issues"
	var args []interface{}

	if status != "" {
		query += " WHERE status = ?"
		args = append(args, status)
	}
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := DB.Query(query, args...)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var issues []Issue
	for rows.Next() {
		var issue Issue
		if err := rows.Scan(&issue.ID, &issue.Title, &issue.Content, &issue.Status); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		issues = append(issues, issue)
	}
	c.JSON(200, issues)
}

func GetIssue(c *gin.Context) {
	id := c.Param("id")
	var issue Issue
	err := DB.QueryRow("SELECT id, title, content, status FROM issues WHERE id = ?", id).
		Scan(&issue.ID, &issue.Title, &issue.Content, &issue.Status)
	if err == sql.ErrNoRows {
		c.JSON(404, gin.H{"error": "not found"})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, issue)
}

func UpdateIssue(c *gin.Context) {
	id := c.Param("id")
	var issue Issue
	if err := c.ShouldBindJSON(&issue); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON"})
		return
	}
	_, err := DB.Exec("UPDATE issues SET title = ?, content = ?, status = ? WHERE id = ?",
		issue.Title, issue.Content, issue.Status, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	issue.ID, _ = strconv.Atoi(id)
	c.JSON(200, issue)
}

func DeleteIssue(c *gin.Context) {
	id := c.Param("id")
	_, err := DB.Exec("DELETE FROM issues WHERE id = ?", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(204)
}
