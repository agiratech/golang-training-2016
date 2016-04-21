package Studentapi
import (
   // "database/sql"
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
var dbmap *gorp.DbMap

func Dbcon(dbm *gorp.DbMap){
	dbmap=dbm
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}


var err error

func GetUsers(c *gin.Context) {
 	var Students []Student

 	_, err = dbmap.Select(&Students, "select * from student_details order by student_id")
    checkErr(err, "Select failed")
 
        c.JSON(200, Students)
   
}

func GetUser(c *gin.Context) {
    var student Student
 	id := c.Param("id")
 	user_id, _ := strconv.ParseInt(id, 10, 64)
 	i:=int32(user_id)
 
  	err = dbmap.SelectOne(&student, "select * from student_details where student_id=$1",i)
 
   // checkErr(err, "Select failed")
    
    if err == nil {
       c.JSON(200, student)
     } else {
       c.JSON(404, gin.H{"error": "user not found"})
     }

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
    //Params like students?id=3&joindate=2016-02-02&name=reddy&address=tirupati
    var student Student
    s:=c.Query("id")
    student.Id,_=strconv.Atoi(s)
    student.JoinDate=c.Query("joindate")
    student.Name=c.Query("name")
    student.Address=c.Query("address")
    err = dbmap.Insert(&student)
   // log.Print(err)
    if err == nil{
        c.String(http.StatusOK,"Inserted Successfully\n")
        c.JSON(200,student)
    } else{
        c.JSON(404,gin.H{"error": "Insertion failed"})
    }
}

func UpdateUser(c *gin.Context){
    //Params like students?id=3&joindate=2016-02-02&name=reddy&address=tirupati
    var student Student
    s:=c.Query("id")
    student.Id,_=strconv.Atoi(s)
    student.JoinDate=c.Query("joindate")
    student.Name=c.Query("name")
    student.Address=c.Query("address")
    _,err = dbmap.Update(&student)
   // log.Print(err)
    if err == nil{
        c.String(http.StatusOK,"Updated Successfully\n")
        c.JSON(200,student)
    } else{
        c.JSON(404,gin.H{"error": "Updation failed"})
    }
}

