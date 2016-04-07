package search
import (
	"log"
	"database/sql"
	_"github.com/lib/pq"
	"encoding/json"
)
type info struct{
	Empid int
	Name string
	Age int
	Address string
	Salary string	
}

func Search(id int) string{
		db,err := sql.Open("postgres","user=postgres password='root' dbname=sai sslmode=disable")
		checkErr(err)

	defer db.Close()
	res,err := db.Query("select * from employee where empid= $1", id)
	checkErr(err)
	var empinfo info
	for res.Next(){
		
		err = res.Scan(&empinfo.Empid,&empinfo.Name,&empinfo.Age,&empinfo.Address,&empinfo.Salary)
		
		}
		js,_ := json.Marshal(empinfo)
		return string(js)

}
func checkErr(err error){
	if err !=nil{
		log.Fatal(err)
	}
}