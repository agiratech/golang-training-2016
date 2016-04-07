package Insertion
import (
	"log"
	"database/sql"
	_"github.com/lib/pq"
	"encoding/json"
	//"fmt"
)
type info struct{
	Empid int `json:",string"`
	Name string
	Age int `json:",string"`
	Address string
	Salary float64 `json:",string"`
}
var result,js1 string
func Insert(empinfo string) (string, string){
		var pro info
		//fmt.Printf("%q\n",empinfo)
		//fmt.Printf("%T %T %T %T %T\n", pro.Empid,pro.Name,pro.Age,pro.Address,pro.Salary)
		err := json.Unmarshal([]byte(empinfo), &pro)
		//fmt.Printf("%T %T %T %T %T\n", pro.Empid,pro.Name,pro.Age,pro.Address,pro.Salary)
		checkErr(err)
		db,err := sql.Open("postgres","user=postgres password='root' dbname=sai sslmode=disable")
		checkErr(err)

	defer db.Close()
	
	stmt,err := db.Prepare("insert into employee(empid,name,age,address,salary)values($1,$2,$3,$4,$5)")
	checkErr(err)
	res1,err :=stmt.Exec(pro.Empid,pro.Name,pro.Age,pro.Address,pro.Salary)
	checkErr(err)
	js,_:= json.Marshal(pro)
		js1=string(js)
	if res1!=nil{
		
		result="Inserted Successfully"
	} else {
		result="Not inserted"
	}
return result,js1
}
func checkErr(err error){
	if err !=nil{
		log.Fatal(err)
	}
}