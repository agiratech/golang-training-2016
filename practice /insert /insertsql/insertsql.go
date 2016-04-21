package main

import ( "fmt"
         "log"
         "sync"
         "time"
         "database/sql"
         //"encoding/json"
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

//func dosomething(milliseconds time.Duration,wg *sync.WaitGroup) {
   
  // duration:=milliseconds * time.Second
   //time.Sleep(duration)
   //fmt.Println("Function in background, Duration:",duration)
   //wg.Done()

//}

func insert(milliseconds time.Duration,wg *sync.WaitGroup) {
    
    duration :=milliseconds * time.Millisecond
    time.Sleep(duration)
   // var wg sync.WaitGroup
    var e info
    var c RecordsInfo
    fmt.Println("enter the id ,name,dep")
    fmt.Scanf("%d",&e.Id)
    fmt.Scanf("%s",&e.Name)
    fmt.Scanf("%s",&e.Dep)
    db,err:= sql.Open("postgres","user=postgres password='root' dbname=aaron sslmode=disable")
    checkerror(err)
    StartTime := time.Now()
    wg.Add(1)
    //dosomething(2,&wg)
    //for c.Records_insert=0 ; c.Records_insert < 100 ; c.Records_insert++ {
    stmt,err := db.Prepare("insert into student values ($1,$2,$3)")
    checkerror(err)
    stmt.Exec(&e.Id,&e.Name,&e.Dep)
    //}
    if c.Records_insert>0 {
        c.Status = "Records Inserted Successfully" 
    }else {
        c.Status = "Records Not Inserted Successfully"
    }
    Stoptime := time.Now()
    d:= Stoptime.Sub(StartTime)
    c.Duration_timeinSec =d.Seconds()
    fmt.Println(e.Id,e.Name,e.Dep)
    wg.Wait()
  //  wg.Done()
   // js,_:=json.Marshal(c)
    //return string(js)

}

func insert1(milliseconds time.Duration,wg *sync.WaitGroup) {
    
    duration :=milliseconds * time.Millisecond
    time.Sleep(duration)
   // var wg sync.WaitGroup
    var e info
    var c RecordsInfo
    fmt.Println("enter the id ,name,dep")
    fmt.Scanf("%d",&e.Id)
    fmt.Scanf("%s",&e.Name)
    fmt.Scanf("%s",&e.Dep)
    db,err:= sql.Open("postgres","user=postgres password='root' dbname=aaron sslmode=disable")
    checkerror(err)
    StartTime := time.Now()
    wg.Add(1)
    //dosomething(2,&wg)
    //for c.Records_insert=0 ; c.Records_insert < 100 ; c.Records_insert++ {
    stmt,err := db.Prepare("insert into student values ($1,$2,$3)")
    checkerror(err)
    stmt.Exec(&e.Id,&e.Name,&e.Dep)
    //}
    if c.Records_insert>0 {
        c.Status = "Records Inserted Successfully" 
    }else {
        c.Status = "Records Not Inserted Successfully"
    }
    Stoptime := time.Now()
    d:= Stoptime.Sub(StartTime)
    c.Duration_timeinSec =d.Seconds()
    fmt.Println(e.Id,e.Name,e.Dep)
    wg.Wait()
    //wg.Done()
   // js,_:=json.Marshal(c)
    //return string(js)

}

func insert2(milliseconds time.Duration,wg *sync.WaitGroup) {
    
    duration :=milliseconds * time.Millisecond
    time.Sleep(duration)
   // var wg sync.WaitGroup
    var e info
    var c RecordsInfo
    fmt.Println("enter the id ,name,dep")
    fmt.Scanf("%d",&e.Id)
    fmt.Scanf("%s",&e.Name)
    fmt.Scanf("%s",&e.Dep)
    db,err:= sql.Open("postgres","user=postgres password='root' dbname=aaron sslmode=disable")
    checkerror(err)
    StartTime := time.Now()
    wg.Add(1)
    //dosomething(2,&wg)
    //for c.Records_insert=0 ; c.Records_insert < 100 ; c.Records_insert++ {
    stmt,err := db.Prepare("insert into student values ($1,$2,$3)")
    checkerror(err)
    stmt.Exec(&e.Id,&e.Name,&e.Dep)
    //}
    if c.Records_insert>0 {
        c.Status = "Records Inserted Successfully" 
    }else {
        c.Status = "Records Not Inserted Successfully"
    }
    Stoptime := time.Now()
    d:= Stoptime.Sub(StartTime)
    c.Duration_timeinSec =d.Seconds()
    fmt.Println(e.Id,e.Name,e.Dep)
    wg.Wait()
    //wg.Done()
   // js,_:=json.Marshal(c)
    //return string(js)

}

func main() {

        var wg sync.WaitGroup
        insert(200,&wg)
        insert1(300,&wg)
        insert2(400,&wg)      
}


