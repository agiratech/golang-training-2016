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
	var sal float64
	fmt.Println("please enter salary:")
	fmt.Scanf("%f",&sal)

	db, err := sql.Open("postgres","user=postgres password='root' dbname=sai sslmode=disable")
	checkErr(err)
	//fmt.Println(string(db.Ping()))
	defer db.Close()
	res,err := db.Query("select * from employee where salary<= $1", sal)
	checkErr(err)

	for res.Next(){
		var empinfo info
		err = res.Scan(&empinfo.Empid,&empinfo.Name,&empinfo.Age,&empinfo.Address,&empinfo.Salary)
		search(empinfo)
		//fmt.Println(&empinfo)
	}
}
func search(empinfo info){
	
		//fmt.Println("EmpID |EmpName	|age|  Address| Salary")
		//fmt.Println(empinfo)
		js,_ := json.Marshal(empinfo)
		//checkErr(err)
		fmt.Println(string(js))


}
func checkErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}