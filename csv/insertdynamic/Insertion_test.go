package Insertion
import (
	"testing"
	"encoding/json"
)
type details struct{
	Empid int `json:",string"`
	Name string
	Age int `json:",string"`
	Address string
	Salary float64  `json:",string"`
}

func Test_Insert(t *testing.T){
	//k:="{Empid:301,Name:sai,Age:22,Address:tirupati,Salary:₹ 10,000.00}"
	var d details
	d.Empid=305
	d.Name="sai"
	d.Age=22
	d.Address="tpt"
	d.Salary=10000
	js,_:= json.Marshal(d)
	js1:=string(js)
	s1,val:= Insert(js1)
	t.Log(val)
if s1=="Inserted Successfully" && val==`{"Empid":"305","Name":"sai","Age":"22","Address":"tpt","Salary":"10000"}`{
	t.Log("Its working")
} else{
	t.Error("Its not working")
}
}