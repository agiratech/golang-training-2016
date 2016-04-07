 package main

import (
 // "bufio"
  "log"
  "os"
  "fmt"
 // "time"
 // "github.com/ajays20078/go-ini"
  //"github.com/gorilla/mux"
  "github.com/ajays20078/go-http-logger"
  "strconv"
  //"./insertdynamic"

  "encoding/json"
  "net/http"
  //"log/syslog"
)
type details struct{
  Empid int `json:",string"`
  Name string
  Age int `json:",string"`
  Address string
  Salary float64  `json:",string"`
}
var d details
// func checkerror(err error) {

//   logfile,err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
//   if err != nil {
//     panic(err)
//   }
//   defer logfile.Close()
//  } 

// func logg() {

//   f,_:=os.Open("log.txt")
//   scanner := bufio.NewReader(f)
//    s,_:= scanner.ReadString(10)
//   fmt.Println(s)
// }

// func profile(w http.ResponseWriter, r *http.Request) {
//   params := mux.Vars(r)
//   name := params["name"]
//   return_string := "Hello " + name
//   ip, port, _ := net.SplitHostPort(r.RemoteAddr)
//   return_string = return_string + "\t IP is" + ip + ":" + port
//   w.Write([]byte(return_string))
//  }

func main() {

  // logwriter, e := syslog.New(syslog.LOG_NOTICE, "myprog")
    //if e == nil {
      //  log.SetOutput(logwriter)
    //}
    // file, err := os.Open("ss.txt")
    //  if err!=nil {
    //   start := time.Now()
    //      checkerror(err)
    //      logg()
    //   stop := time.Now()
    //  fmt.Println(start,stop)
    //     }
  logprofile,err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
  if err != nil {
    panic(err)
  }
  // defer logfile.Close()
    
  //   scanner := bufio.NewScanner(file)
  //   for scanner.Scan() {
  //       log.Println(scanner.Text())
  //   }
  http.Handle("/",http.FileServer(http.Dir("index.html")))
  http.HandleFunc("/empid",handler)
    //rtr := mux.NewRouter()
  //rtr.HandleFunc("/user/{name:[a-z]+}/profile", profile).Host(virtual_server)
  //http.Handle("/", rtr)
  log.Println("Listening...")
  http.ListenAndServe(":1000", httpLogger.WriteLog(http.DefaultServeMux,logprofile))

}
func handler(w http.ResponseWriter,r *http.Request){
  id :=r.FormValue("empid")
  d.Empid,_= strconv.Atoi(id)
  d.Name =r.FormValue("name")
  age :=r.FormValue("age")
  d.Age,_= strconv.Atoi(age)
  d.Address =r.FormValue("address")
  salary :=r.FormValue("salary")
  d.Salary,_= strconv.ParseFloat(salary, 64)
  js,_ := json.Marshal(d)
  js1:=string(js)
  result,values :=Insertion.Insert(js1)
  fmt.Fprintf(w, "Status : %s \n%s!",values,result)
   // ip, port, _ := net.SplitHostPort(r.RemoteAddr)


}
// http.Handle("/", http.FileServer(http.Dir("static")))
//   err := http.ListenAndServe(":3002", Log(http.DefaultServeMux))
//   if err != nil {
//     log.Fatal("ListenAndServe: ", err)
//   }
// }

// func Log(handler http.Handler) http.Handler {
//   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     fmt.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.Proto)
//     handler.ServeHTTP(w, r)
//   })
// }