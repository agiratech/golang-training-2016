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
	d.Empid=303
	d.Name="sai"
	d.Age=22
	d.Address="tpt"
	d.Salary=10000
	js,_:= json.Marshal(d)
	js1:=string(js)
	s1,_:= Insert(js1)
if s1=="Inserted Successfully"{
	t.Log("Its working")
} else{
	t.Error("Its not working")
}
}