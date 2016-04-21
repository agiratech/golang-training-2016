package main 

import (
"./insertsql"
"net/http"
)

var in string	

func main(){
	
	in = insertsql.Db_con()
	http.Handle("/",http.FileServer(http.Dir("insert")))
	http.HandleFunc("/insertsql",handler)
	http.ListenAndServe(":5011",nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(in))
}