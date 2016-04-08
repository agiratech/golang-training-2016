package Errorpac
import(
	"log"
	"os"
)

func CheckErr(err error) {
	if err!=nil{
	 Errorfile()
	}
}
func Errorfile(){
	logfile,err := os.OpenFile("log.log", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
  	log.SetOutput(logfile)
  
  	log.Println(err)
  

	//defer logfile.Close()
}