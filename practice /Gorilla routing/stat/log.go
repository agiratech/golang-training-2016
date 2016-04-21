package main 

import ( "log" 
         "os")

func main() {
	
	logfile,err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
	if err!=nil{
	 panic(err)
    }
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Println("some error in your file")
}
