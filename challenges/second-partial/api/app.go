package main

import (
  "api/src/routes"
  "net/http"
  "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
    app := gin.Default()
    app.MaxMultipartMemory = 8 << 20 // 8 MiB

    // ******
    // PUBLIC
    // ******
    app.LoadHTMLGlob("public/*")

    app.GET("/", func(c *gin.Context) {
      c.HTML(
          http.StatusOK,
          "index.html",
          gin.H{
              "title": "Home Page",
          },
      )
    })

    // ******
    // ROUTES
    // ******
    routes.Init(app)

    app.Run(":8080")
}
