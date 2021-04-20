package controllers

import (
  "api/src/models"
  "api/src/utils"
  "fmt"
  "time"
  "strings"
  "net/http"
	"github.com/gin-gonic/gin"
)

// Login
func Login(c *gin.Context) {

  var body utils.Authentication
  c.BindJSON(&body)

	val, err := models.Login(&body)

	if err != nil {
		fmt.Println(err.Error())
    c.JSON(http.StatusNotFound, err.Error())
	} else {
		c.JSON(http.StatusOK, val)
	}

}

// Logout
func Logout(c *gin.Context) {

  user, _ := c.Get("User")
  account := user.(*utils.ClaimsStruct)

  token, _ := c.Get("Token")
  jwtToken, _ := token.(string)

  // Revoke Token
  utils.Tokens = utils.Remove(utils.Tokens, jwtToken)

  val := utils.MessageStruct {
    Message: "Bye " + account.User + ", your token has been revoked",
  }

	c.JSON(http.StatusOK, val)
}

// Status
func Status(c *gin.Context) {
  user, _ := c.Get("User")
  account := user.(*utils.ClaimsStruct)

  currTime := time.Now().String()
  splitTime := strings.Split(currTime, ".")
  time := splitTime[0]

  val := utils.MessageStruct {
    Message: "Hi " + account.User + ", the DPIP System is Up and Running",
    Time: time,
  }

  c.JSON(http.StatusOK, val)

}
