package main
import(
	// "fmt"
	"net/http"
	//"./saipackage"
	//"./insert"
	"github.com/agiratech/golang-training-2016/reddysai/insert"
)
var js string
func main() {
	//js = saipackage.Searchkey()

	js =insert.Insertion()

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/insert", handler)

	http.ListenAndServe(":5500", nil)
	//fmt.Println("------")
}

func handler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(js))
}
 //"github.com:agiratech/golang-training-2016/reddysai"