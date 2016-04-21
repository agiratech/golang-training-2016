 package main

 import (
         "net/http"
         "github.com/gin-gonic/gin"
 )

 func main() {
 	r := gin.Default()

 	r.GET("/query", func(c *gin.Context) {
 		first := c.Request.URL.Query().Get("first") //<---- here!
 		c.String(http.StatusOK, "First is "+first+"\n")

 		second := c.Request.URL.Query().Get("second")
 		c.String(http.StatusOK, "Second is "+second+"\n")
 	})

 	r.Run(":8080") // listen and serve on 0.0.0.0:8080

 }