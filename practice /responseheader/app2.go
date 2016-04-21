package main
import (
	"log"
	"net/http"
	"fmt"

)

func main() {
	http.Handle("/",http.FileServer(http.Dir("stat")))
	http.HandleFunc("http://localhost:3001/arun",handler)
	log.Fatal(http.ListenAndServe(":4001",nil))
}
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "hello, ")
}


        client := &http.Client{}

        req, err := http.NewRequest("GET", "http://httpbin.org/user-agent", nil)
        if err != nil {
                log.Fatalln(err)
        }

        req.Header.Set("User-Agent", "Golang Spider Bot v. 3.0")

        resp, err := client.Do(req)
        if err != nil {
                log.Fatalln(err)
        }

        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                log.Fatalln(err)
        }

        log.Println(string(body))

}