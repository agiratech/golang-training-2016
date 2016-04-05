package insert
import (
 // "os"
  "log"
  "time"
  //"fmt"
  "database/sql"
  _ "github.com/lib/pq"
  	"encoding/json"

)
type RecordInfo struct{
	Records int
	Status string
	Duration_sec float64
}
var db *sql.DB
var err error
var id,i int
var info RecordInfo
var r int=0
func DBcon() {
	db,err = sql.Open("postgres","user=postgres password='root' dbname=sai sslmode=disable")
	checkErr(err)
	//defer db.Close()
}
func Insertion() string {
	t1 := time.Now()
	for i=0;i<20;i++{
		DBcon()
	res,err := db.Query("select max(empid) from employee")
	
	
	checkErr(err)
		for res.Next(){
		res.Scan(&id)
		}
	//fmt.Println(id)
	stmt,err := db.Prepare("insert into employee(empid,name,age,address,salary)values($1,$2,$3,$4,$5)")
	checkErr(err)
	res1,err :=stmt.Exec(id+1,"malli",24,"Dharmavaram",10000)
	checkErr(err)
	if res1!=nil{
		r=1
	}
		
	}
	t2 := time.Now()
		if r==1{
			info.Records=i
			t:=t2.Sub(t1)
			info.Duration_sec= t.Seconds()
			info.Status="Inserted Sucessfully"
		}
		js,_ := json.Marshal(info)
		defer db.Close()
	return string(js)

}
func checkErr(err error){
	if err !=nil{
	log.Fatal(err)
	}
}

func main(){
	DBcon()
	Insertion()
	
}