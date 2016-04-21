package main

import (
    "database/sql"
    "gopkg.in/gorp.v1"
   // _ "github.com/mattn/go-sqlite3"
    "log"
   // "time"
    "github.com/gin-gonic/gin"
    "strconv"
     _"github.com/lib/pq"
     "net/http"
)

type Student struct {
    
    Id      int  `db:"student_id" json:"student_id"`
    JoinDate string
    Name   string `db:"student_name"`              
    Address    string `db:"student_address"` 
}

// func newPost(name, address string) Student {
//     return Student{
//         JoinDate: time.Now().UnixNano(),
//         Name:   name,
//         Address:    address,
//     }
// }

func initDb() *gorp.DbMap {
    
    db, err := sql.Open("postgres", "user=postgres password='root' dbname=sai sslmode=disable")
    checkErr(err, "sql.Open failed")

    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

    dbmap.AddTableWithName(Student{}, "student_details").SetKeys(true, "Id")

    err = dbmap.CreateTablesIfNotExists()
    checkErr(err, "Create tables failed")

    return dbmap
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}
var dbmap *gorp.DbMap
var err error
func main() {
 r := gin.Default()
 dbmap = initDb()

 r.GET("/users", GetUsers)
 r.GET("/users/:id", GetUser)
 r.DELETE("/users/:id",DeleteUser)
 r.POST("/users", PutUser)
 //v1.PUT("/users/:id", UpdateUser)
 //v1.DELETE("/users/:id", DeleteUser)
 //}

r.Run(":8080")
}
func GetUsers(c *gin.Context) {
 var Students []Student

 _, err = dbmap.Select(&Students, "select * from student_details order by student_id")
    checkErr(err, "Select failed")
    //c.String("All rows:")
    //for _, _ := range Students {
        c.JSON(200, Students)
   // }

// curl -i http://localhost:8080/api/v1/users
}
func GetUser(c *gin.Context) {
    var student Student
 id := c.Param("id")
 user_id, _ := strconv.ParseInt(id, 10, 64)
 i:=int32(user_id)
 
  err = dbmap.SelectOne(&student, "select * from student_details where student_id=$1",i)
 
    checkErr(err, "Select failed")
    
             if err == nil {
                 
             c.JSON(200, student)
             } else {
             c.JSON(404, gin.H{"error": "user not found"})
             }
    //c.JSON(200, student)
 }
 func DeleteUser(c *gin.Context){
    id := c.Param("id")
    user_id, _ := strconv.ParseInt(id, 10, 64)
    i:=int32(user_id)

var student Student
 err = dbmap.SelectOne(&student, "SELECT student_id FROM student_details WHERE student_id=$1", i)

if err == nil {
 _, err = dbmap.Delete(&student)

    if err == nil {
     c.JSON(200, gin.H{"id #" + id: " deleted"})
     } else {
     checkErr(err, "Delete failed")
     }

} else {
 c.JSON(404, gin.H{"error": "user not found"})
 }
 }
 func PutUser(c *gin.Context){
    //Params like insert?id=3&joindate=2016-02-02&name=reddy&address=tirupati
    var student Student
    s:=c.Query("id")
    student.Id,_=strconv.Atoi(s)
    student.JoinDate=c.Query("joindate")
    student.Name=c.Query("name")
    student.Address=c.Query("address")
    err = dbmap.Insert(&student)
    log.Print(err)
    if err == nil{
        c.String(http.StatusOK,"Inserted Successfully\n")
        c.JSON(200,student)
    } else{
        c.JSON(404,gin.H{"error": "Insertion failed"})
    }
 }


