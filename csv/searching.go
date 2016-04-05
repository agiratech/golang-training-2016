package main


import (
		"fmt"
		// "encoding/csv"
		//  "os"
		   "log"
  		//	"time"
  		"database/sql"
  _ "github.com/lib/pq"
  )
type info struct{
	empid int
	name string
	age int
	address string
	salary string
}

  func main() {
  		db, err := sql.Open("postgres","user=postgres password='root' dbname=sai sslmode=disable")
  		checkErr(err)
  		defer db.Close()

  		var id int
  		fmt.Println("please enter Employee id :")
  		fmt.Scanf("%d",&id)
  		//res1:= fmt.Sprintf("select * from employee where empid=$1",id)
  		//query, err := db.Prepare("select * from employee where empid=$1")
  		//checkErr(err)
  		//res1, err := query.Exec(id)
  		//checkErr(err)
  		res, err := db.Query("select * from employee where empid= $1", id)
  		checkErr(err)

	  		for res.Next() {

		  		/*var empid int
		  		var name string
		  		var age int
		  		var address string
		  		var salary string*/
		  		var idinfo info
		  		err = res.Scan(&idinfo.empid,&idinfo.name,&idinfo.age,&idinfo.address,&idinfo.salary)
		  		//err = res.Scan(&empid,&name,&age,&address,&salary)
		  		checkErr(err)
		  		fmt.Println("EmpID |EmpName	|age|  Address| Salary")
		  		//fmt.Println(empid,"\t",name,"\t",age,"\t",address," ",salary)
		  		fmt.Println(idinfo)
	  		}
  }
  func checkErr(err error) {
  		if err !=nil{

  		log.Fatal(err)
  		}
  }