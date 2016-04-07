package main
import(
	"encoding/json"
	"fmt"
)
type details struct{
	Empid int //`json:",string"`
	Name string
	Age int //`json:",string"`
	Address string
	Salary float64 `json:",string"`
}
func main() {
		s := `{"Empid":4,"Name":"adil","Age":23,"Address":"Nellore","Salary":"20000.00"}`
    var pro details
    err := json.Unmarshal([]byte(s), &pro)
    if err == nil {
        fmt.Printf("%T %T %T %T %T\n", pro.Empid,pro.Name,pro.Age,pro.Address,pro.Salary)
    } else {
        fmt.Println(err)
        fmt.Printf("%+v\n", pro)
    }
	
}