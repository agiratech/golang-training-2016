 package main 
import(
     "github.com/gin-gonic/gin"
     //"net/http"
 )

type details struct{
    Firstname string
    Lastname string
}
func main() {
    router := gin.Default()
    var d details
    // Query string parameters are parsed using the existing underlying request object.
    // The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
    router.GET("/welcome", func(c *gin.Context) {
        d.Firstname= c.DefaultQuery("firstname", "Guest")
        d.Lastname= c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
       // d.Firstname=firstname
       // d.Lastname=lastname
       // c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
        c.JSON(200,d)
    })
    router.Run(":8000")
}