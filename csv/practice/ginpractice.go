// package main 
// import(
//     "github.com/gin-gonic/gin"
//     "net/http"
// )
// func main() {
//     router := gin.Default()

//     // This handler will match /user/john but will not match neither /user/ or /user
//     router.GET("/user/:name", func(c *gin.Context) {
//         name := c.Param("name")
//         c.String(http.StatusOK, "Hello %s", name)
//     })

//     // However, this one will match /user/john/ and also /user/john/send
//     // If no other routers match /user/john, it will redirect to /user/john/
//     router.GET("/user/:name/*action", func(c *gin.Context) {
//         name := c.Param("name")
//         action := c.Param("action")
//         message := name + " is " + action
//         c.String(http.StatusOK, message)
//     })

//     router.Run(":8080")
// }
package main

import (
    "github.com/gin-gonic/gin"
)

func index (c *gin.Context){
    content := gin.H{"url" : "http://www.flipkart.com/search?q=samsung%20mobiles&as=on&as-show=on&otracker=start&as-pos=3_q_s" }
    c.JSON(200, content)
}

func main(){
  app := gin.Default()
  app.GET("/", index)
  app.Run(":8000")
}
