package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func index (c *gin.Context){
   // content := gin.H{"url": "[https://www.google.co.in/webhp?hl=en]"}
 
    r, _ := http.Get("https://www.google.co.in/webhp?hl=en")
    c.JSON(200, r)
    //defer r.Body.Close()
   // dec := json.NewDecoder(r.Body)
}

func main(){
  app := gin.Default()
  app.GET("/", index)
  app.Run(":8080")
}

 