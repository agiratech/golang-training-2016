package main

import (
    //"fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "strconv"
    "./searchdynamic"
    "os"
    "time"
    "path"
   // "encoding/json"
)
type LogRecord struct{
    http.ResponseWriter
    status int
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
   // fmt.Fprint(w, "Welcome!\n")
    id :=r.FormValue("empid")
    s,_:= strconv.Atoi(id)

    js := search.Search(s)
    w.Write([]byte(js))
}
func (r *LogRecord) Write(p []byte) (int,error){
    return r.ResponseWriter.Write(p)
}

func (r *LogRecord) WriteHeader(status int){
    r.status =status
    r.ResponseWriter.WriteHeader(status)
}
func WrapHandler(f http.Handler,logfile *os.File) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request){
        logger := log.New(logfile, "", 0)
        writer := LogRecord{w, 0}
        //byte, _ := json.Marshal(r)
        r.Header.Set("Content-Type","application/json")
        //byte, _ := json.Marshal(r.Header.Get("Accept"))
        //log.Printf("%v",string(byte))
        f.ServeHTTP(&writer, r)
        statusCode := writer.status
        end:=time.Now()
        s:="==========================================="
      // log.Printf("\n %v\n %v\n %v\n",r.Form,r.Header,w)
       // r.Header.Get(Content-Type)
        log.Printf("%sHost : %s Remote Address : %s  Method Type : %s  Path : %s  Prote Type : %s  StatusCode : %d  Content Type : %v",s,r.Host, r.RemoteAddr, r.Method, r.URL.Path, r.Proto,statusCode,r.Header.Get("Accept"))
        logger.Printf("%s  Time : %v  Host : %s  Remote Address : %s  Method Type : %s  Path : %s  Prote Type : %s  StatusCode : %d  HttpClient : %v  Content Type : %v",s,end,r.Host, r.RemoteAddr, r.Method, r.URL.Path, r.Proto,statusCode,r.Header.Get("User-Agent"),r.Header.Get("Accept"))

    }
}


func main() {
    router1 := httprouter.New()
   // router :=http.NewServeMux()
    //Its serves static files
     router1.NotFound = http.FileServer(http.Dir("static/"))
     router1.GET("/empid", Index)
     log.Printf("Direction %s",path.Clean("static"))
     logfile,_ := os.OpenFile("error/log.log", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
    log.Fatal(http.ListenAndServe(":8080", WrapHandler(router1,logfile)))
}


