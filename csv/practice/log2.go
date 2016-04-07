package main


 import (
 	"flag"
 	"github.com/ajays20078/go-ini"
 	"github.com/gorilla/mux"
 	"github.com/ajays20078/go-http-logger"
 	"log"
 	"net"
 	"net/http"
 	"os"
 	
 )
 
 func check(e error) {
 	if e != nil {
 		panic(e)
 	}
 }
 
 func check_boolean(e bool) {
 	if !e {
 		panic("Unexpected Error!!!")
 	}
 }
 
 func main() {
 
 	wordPtr := flag.String("c", "/etc/goserver/config.ini", "config file")
 	file, err := ini.LoadFile(*wordPtr)
 	check(err)
 	port, ok := file.Get("Server", "Port")
 	check_boolean(ok)
 	virtual_server, ok := file.Get("Server", "VirtualHost")
 	check_boolean(ok)
 	access_log, ok := file.Get("Server", "AccessLog")
 	check_boolean(ok)
 	access_file_handler, err := os.OpenFile(access_log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
 	if err != nil {
 		panic(err)
 	}
 	defer access_file_handler.Close()
 	rtr := mux.NewRouter()
 	rtr.HandleFunc("/user/{name:[a-z]+}/profile", profile).Host(virtual_server)
 	http.Handle("/", rtr)
 	log.Println("Listening...")
 	http.ListenAndServe(":"+port, httpLogger.WriteLog(http.DefaultServeMux,access_file_handler))
 }
 
 func profile(w http.ResponseWriter, r *http.Request) {
 	params := mux.Vars(r)
 	name := params["name"]
 	return_string := "Hello " + name
 	ip, port, _ := net.SplitHostPort(r.RemoteAddr)
 	return_string = return_string + "\t IP is" + ip + ":" + port
 	w.Write([]byte(return_string))
 }
 