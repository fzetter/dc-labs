package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// About
func About(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"class": "Distributed Computing",
			"student": "Fernanda Zetter",
			"challenge": "second-partial",
	})
}
