package main

import ("database/sql"
        "github.com/gin-gonic/gin"
        _"github.com/lib/pq"
        "gopkg.in/gorp.v1"
        "fmt"
        "log"
        "strconv")

  

type Invoice struct {
    Id     int64 `db:"id" json:"id"`
    Created  int64 `db:"created" json:"created"`
    Updated  int64 `db:"updated" json:"updated"`
    Memo     string `db:"memo" json:"memo"`
    PersonId int64  `db:"personId" json:"personId"`

}

//var dbmap = initDb()
var dbmap *gorp.DbMap

func checkerr(err error, msg string){
 
	if err!=nil {
	log.Fatal(err,msg)
	}
}

func initDb() *gorp.DbMap {
	 
	 db, err := sql.Open("postgres","user=postgres password='root' dbname=aaron sslmode=disable")
     checkerr(err,"sql.Open failed")    
	 dbmap := &gorp.DbMap{Db: db, Dialect:gorp.PostgresDialect{}}

     return dbmap 
}

func createInvoice(c *gin.Context) {
    

  //var i string
  //fmt.Println("Enter the Table Name You Want to Create")
  //fmt.Scanf("%s",&i)
    dbmap.AddTableWithName(Invoice{},"invoice").SetKeys(true, "Id")
    err := dbmap.CreateTablesIfNotExists()
    checkerr(err,"Create table failed")

}

func getInvoice(c *gin.Context) {
   
   //dbmap:=initDb
   var invoice []Invoice
   _,err := dbmap.Select(&invoice,"select * from arun")
   if err == nil {
   	c.JSON(200,invoice)
   } else {
   	c.JSON(404,gin.H{"error":"no Inovice(s) into the table"})
   }

}

func getInvoiceid(c *gin.Context) {
   
   //var i Invoice 
   id:= c.Param("id")
   zod,_:=strconv.ParseInt(id, 10, 64)
   var invoice []Invoice
   _,err := dbmap.Select(&invoice,"select * from arun where id=$1",zod)
   if err == nil {
   	c.JSON(200,invoice)
   } else {
   	c.JSON(404,gin.H{"error":"no Inovice(s) into the table"})
   }
}

var err error
func deleteInvoice(c *gin.Context) {
      
        id := c.Param("id")
     zod,_:=strconv.ParseInt(id, 10, 64)
     //i := int64(zod)
     var invoice []Invoice
     err = dbmap.SelectOne(&invoice,"select id from arun where id=$1",zod)
     log.Print(err)
     if err ==nil{
     	_,err=dbmap.Delete(&invoice)
     	log.Print(err)
     	if err==nil{
     		c.JSON(200,gin.H{"id ="+id:"deleted"})
     	} else{
     		c.JSON(404,gin.H{"error":"not deleted"})
     	}
     } else {
     	c.JSON(404,gin.H{"error":"user not found"})
     }
}

func insertInvoiceid(c *gin.Context) {
   
    var i Invoice 
    fmt.Println("Enter the details")
    fmt.Scanf("%d",&i.Id)
    fmt.Scanf("%d",&i.Created)
    fmt.Scanf("%d",&i.Updated)
    fmt.Scanf("%s",&i.Memo)
    fmt.Scanf("%d",&i.PersonId)
    _,err := dbmap.Exec("insert into arun values($1,$2,$3,$4,$5)",i.Id,i.Created,i.Updated,i.Memo,i.PersonId)
    if err == nil {
    	c.JSON(200,i)
    }else {
    checkerr(err,"Insert isn't Done Properly")
   }
}


func main() {
    
    dbmap = initDb()
    router := gin.Default()
    router.GET("/user1",createInvoice)
    router.GET("/user/",getInvoice)
    router.GET("/user/:id",getInvoiceid)
    router.GET("/insert",insertInvoiceid)
    router.GET("/delete/:id",deleteInvoice)
    router.Run(":4444")
	defer dbmap.Db.Close()
}


