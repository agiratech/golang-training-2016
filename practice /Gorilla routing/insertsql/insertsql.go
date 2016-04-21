package insertsql

import ( "fmt"
         "log"
         "time"
         "database/sql"
         "encoding/json"
         _"github.com/lib/pq"
)

type info struct{
    
    Id int
    Name string
    Dep string
}
type RecordsInfo struct{
    Records_insert int
    Status string
    Duration_timeinSec float64
}

func checkerror(c error){
    if c!=nil{
    log.Fatal(c)
    }
}

func Db_conn() string{
    var e info
    var c RecordsInfo
    fmt.Println("enter the id ,name,dep")
    fmt.Scanf("%d",&e.Id)
    fmt.Scanf("%s",&e.Name)
    fmt.Scanf("%s",&e.Dep)
    db,err:= sql.Open("postgres","user=postgres password='root' dbname=aaron sslmode=disable")
    checkerror(err)
    StartTime := time.Now()
    for c.Records_insert=0 ; c.Records_insert < 100 ; c.Records_insert++ {
    stmt,err := db.Prepare("insert into student values ($1,$2,$3)")
    checkerror(err)
    stmt.Exec(&e.Id,&e.Name,&e.Dep)
    }
    if c.Records_insert>0 {
        c.Status = "Records Inserted Successfully" 
    }else {
        c.Status = "Records Not Inserted Successfully"
    }
    Stoptime := time.Now()
    d:= Stoptime.Sub(StartTime)
    c.Duration_timeinSec =d.Seconds()
    fmt.Println(e.Id,e.Name,e.Dep)
    js,_:=json.Marshal(c)
    return string(js)
}


