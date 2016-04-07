package main
import (
	"fmt"
	"log"
	//"html"
	"net/http"
	//"golang-training-2016/csv/searchdynamic"
	"strconv"
	"./insertdynamic"
	"encoding/json"
	//"time"
)
type details struct{
	Empid int `json:",string"`
	Name string
	Age int `json:",string"`
	Address string
	Salary float64  `json:",string"`
}
var d details
func main() {
	
	http.Handle("/",http.FileServer(http.Dir("practice")))
	http.HandleFunc("/empidm",handler)
	log.Fatal(http.ListenAndServe(":3000",Log(http.DefaultServeMux)))
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

}
func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s %s %s",r.Host, r.RemoteAddr, r.Method, r.URL.Path[1:], r.Proto)
		handler.ServeHTTP(w, r)
})
}
