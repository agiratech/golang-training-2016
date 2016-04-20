package main 
import(
    "github.com/gin-gonic/gin"
    "database/sql"
    "log"
    //"github.com/jmoiron/sqlx"
    _"github.com/lib/pq"
)
type details struct{
    Empid int    `db:"empid" json:"empid"`
    Name string  `db:"name" json:"name"`
    Age int   `db:"age" json:"age"`
    Address string  `db:"address" json:"address"`
    Salary string `db:"salary" json:"salary"`
}
var results details

func SomeHandler(db *sql.DB) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        // Your handler code goes in here - e.g.
        rows, err := db.Query("select * from employee where empid= $1",1)
         checkErr(err)
         for rows.Next(){
            err:=rows.Scan(&results.Empid,&results.Name,&results.Age,&results.Address,&results.Salary)
            checkErr(err)
         }

        c.JSON(200, results)
    }

    return gin.HandlerFunc(fn)
}

func main() {
    db, err := sql.Open("postgres","user=postgres password='root' dbname=sai sslmode=disable")
    checkErr(err)

    router := gin.Default()
    router.GET("/test", SomeHandler(db))
    router.Run(":8080")
}
func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}