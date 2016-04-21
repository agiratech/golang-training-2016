package main

import ( "fmt"
         "log"
         //"sync"
         "time"
         "database/sql"
        _"github.com/lib/pq"
)


const (
  gophers = 10
  entries = 100
)

type info struct {
    
    Id int
    Name string
    Dep string

}
type RecordsInfo struct {

    Records_insert int
    Status string
    Duration_timeinSec float64

}

func checkerror(c error) {

    if c!=nil{
    log.Fatal(c)
    }

}

/* func dosomething(milliseconds time.Duration,wg *sync.WaitGroup) {
   
  duration:=milliseconds * time.Millisecond
  time.Sleep(duration)
  fmt.Println("Function in background, Duration:",duration)
  wg.Done()

} */

func insert(jobQueue chan string) {
    
   // var wg sync.WaitGroup
   // taskQueue := make(chan chan int)
    //defer close(taskQueue)
    var e info
    var c RecordsInfo
    e.Id = 1
    e.Name = "aaron"
    e.Dep = "MCA"
    db,err:= sql.Open("postgres","user=postgres password='root' dbname=aaron sslmode=disable")
    checkerror(err)
    StartTime := time.Now()
    //wg.Add(1)
   // dosomething(200,&wg)
    fmt.Printf("StartTime: %v\n",time.Now())
    for c.Records_insert=0 ; c.Records_insert < gophers ; c.Records_insert++ {
    for i := 0; i<entries; i++ {
    stmt,err := db.Prepare("insert into student values ($1,$2,$3)")
    checkerror(err)
    stmt.Exec(&e.Id,&e.Name,&e.Dep)
    }
}
    if c.Records_insert>0 {
        c.Status = "Records Inserted Successfully" 
    }else {
        c.Status = "Records Not Inserted Successfully"
    }
    Stoptime := time.Now()
    d:= Stoptime.Sub(StartTime)
    c.Duration_timeinSec =d.Seconds()
   // jobQueue <- taskQueue
    jobQueue <- fmt.Sprintf(string(e.Id),e.Name,e.Dep)
    
    fmt.Printf("StopTime: %v\n", time.Now())
   // wg.Wait()
    //wg.Done()
}


func main() {
       
       jobQueue := make(chan string)
       defer close(jobQueue)
       fmt.Println("Begin: Kicking off workers that want to use the job queue")
       go insert(jobQueue)
      // go insert(jobQueue)
       //go insert(jobQueue)      
       remainingJobs := 1
       for job := range jobQueue {
       for task := range job {
       fmt.Printf("Handling: %s\n", task)
       }
       remainingJobs--
       if remainingJobs == 0 {
       break
       }
    }
}


