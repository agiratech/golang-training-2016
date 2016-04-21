package main
import(
"github.com/gin-gonic/gin"
"./gin-api"
"log"
"database/sql"
"gopkg.in/gorp.v1"
)



var dbmap *gorp.DbMap
var err error


func main() {
 r := gin.Default()
 dbmap = initDb()
 Studentapi.Dbcon(dbmap)
 
 r.GET("/students", Studentapi.GetUsers)
 r.GET("/student/:id", Studentapi.GetUser)
 r.DELETE("/student/:id",Studentapi.DeleteUser)
 r.POST("/students", Studentapi.PutUser)
 r.PUT("/students", Studentapi.UpdateUser)
 

r.Run(":8000")
}

func initDb() *gorp.DbMap {
    
    db, err := sql.Open("postgres", "user=postgres password='root' dbname=sai sslmode=disable")
    checkErr(err, "sql.Open failed")

    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

    dbmap.AddTableWithName(Studentapi.Student{}, "student_details").SetKeys(true, "Id")

    err = dbmap.CreateTablesIfNotExists()
    checkErr(err, "Create tables failed")

    return dbmap
}
func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}