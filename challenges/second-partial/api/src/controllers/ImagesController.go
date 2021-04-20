package controllers

import (
  "strconv"
  "net/http"
	"github.com/gin-gonic/gin"
)

// Upload
func Upload(c *gin.Context) {

  // Single File Upload
  file, err := c.FormFile("data")

  // No File Received
  if err != nil {
      c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No File Received"})
      return
  }

  // Resolve
  c.JSON(http.StatusOK, gin.H{
      "message": "An image has been successfully uploaded",
      "filename": file.Filename,
      "size": strconv.FormatInt(int64(file.Size/1000), 10) + "kb",
  })

}
