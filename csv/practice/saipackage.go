package main 

import (
	"fmt"
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
func main() {
	searchkey()
}
func searchkey() {
	var sal float64
	fmt.Println("Enter the salary")
	fmt.Scanf("%f",&sal)

	db,err := sql.Open("postgres","user=postgres password='root' dbname=sai sslmode=disable")
	checkErr(err)
	defer db.Close()
	res,err :=db.Query("select * from employee where salary = $1", sal)
	checkErr(err)
	for res.Next() {
			var empinfo info
		err = res.Scan(&empinfo.Empid,&empinfo.Name,&empinfo.Age,&empinfo.Address,&empinfo.Salary)
		js,_ := json.Marshal(empinfo)
		//checkErr(err)
		fmt.Println(string(js))
	}

}

func checkErr(err error) {
	if err!=nil{
	log.Fatal(err)
	}
}