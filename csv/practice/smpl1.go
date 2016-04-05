package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"database/sql"
)
var index_html []byte
var db *sql.DB
var HOST="localhost"
func main() {
	fmt.Println("starting server on http://localhost:8880/")
	http.HandleFunc("/",IndexHandler)
	http.HandleFunc("/redirect", CreateRedirectHandler)
	http.ListenAndServe(":8880",nil)
}
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Host == HOST {
	log.Println("GET/")
	w.Write([]byte(index_html))
	return	
	}
	RedirectHandler(w,r)
}
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	
}
func CreateRedirectHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("creating direct", r.Form)
}
func init() {
	index_html,_ = ioutil.ReadFile("./static/firstone.html")
	connectToDB()
}
func connectToDB() {
	var err error 
	db, err = sql.Open("postgres","user=postgres password='root' dbname=sai sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	salary :=10000
	rows,err := db.Query("select * from employee where salary = $1",salary)

}