package main
import (
	"fmt"
	"log"
	//"html"
	"net/http"
	"golang-training-2016/csv/searchdynamic"
	"strconv"
)

func main() {
	
	http.Handle("/",http.FileServer(http.Dir("static")))
	http.HandleFunc("http://localhost:8080/empid",handler)
	log.Fatal(http.ListenAndServe(":8000",nil))
}
func handler(w http.ResponseWriter,r *http.Request){
	id :=r.FormValue("empid")
	s,_:= strconv.Atoi(id)

	js := search.Search(s)
	fmt.Fprintf(w,"hello %s", s)
	w.Write([]byte(js))
	
}