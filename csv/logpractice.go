package main
import (
	"fmt"
	"log"
	//"html"
	"net/http"
	//"golang-training-2016/csv/searchdynamic"
	"strconv"
	"./insertdynamic"
	"./error"
	"encoding/json"
	//"time"
	"bytes"
	"os"
)
type details struct{
	Empid int `json:",string"`
	Name string
	Age int `json:",string"`
	Address string
	Salary float64  `json:",string"`
}
var d details
var buf bytes.Buffer
var err error
func main() {
	
	http.Handle("/",http.FileServer(http.Dir("practice")))
	http.HandleFunc("/empid",handler)
	log.Fatal(http.ListenAndServe(":3000",Log(http.DefaultServeMux)))
}
func handler(w http.ResponseWriter,r *http.Request){
	
	id :=r.FormValue("empid")
	d.Empid,err= strconv.Atoi(id)
	Errorpac.CheckErr(err)
	d.Name =r.FormValue("name")
	age :=r.FormValue("age")
	d.Age,err= strconv.Atoi(age)
	Errorpac.CheckErr(err)
	d.Address =r.FormValue("address")
	salary :=r.FormValue("salary")
	d.Salary,err= strconv.ParseFloat(salary, 64)
	Errorpac.CheckErr(err)
	js,err := json.Marshal(d)
	Errorpac.CheckErr(err)
	js1:=string(js)
	result,values :=Insertion.Insert(js1)
	fmt.Fprintf(w, "Status : %s \n%s!",values,result)

	

}
func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s %s %s",r.Host, r.RemoteAddr, r.Method, r.URL.Path, r.Proto)
		handler.ServeHTTP(w, r)
		 logfile,_ := os.OpenFile("error/log.log", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
  		 log.SetOutput(logfile)
  		 
 		 //log.Println(err)


		})
}

