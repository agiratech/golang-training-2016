package stringutil
import (
	"testing"
	"github.com/stretchr/testify/assert"
	   . "gopkg.in/check.v1"
)


func Test_StrLength(t *testing.T){
	var s string="sai"
	// if StrLength(s)!=3{
	// 	t.Error("its not working")
	// } else {
	// 	t.Log("its working")
	// }
	assert.Equal(t, 3, StrLength(s), "The two words should be the same.")
}
func Test_StrconvItoa(t *testing.T){
	
	if StrconvItoa(3)!="3"{
		t.Error("Its not working")
	} else {
		t.Log("its working")
	}
}
func Test_StrconvAtoi(t *testing.T){
	if StrconvAtoi("3")!=3{
		t.Error("Its not working")
	} else {
		t.Log("its working")
	}
}
func Test_StrconvItof(t *testing.T){
	if StrconvItof(3)!=3.0000{
		t.Error("Its not working")
	} else {
		t.Log("its working")
	}
}
func Test_StrconvFtoi(t *testing.T){
	var i int
	i=StrconvFtoi(3.14)
	if i!=3{
		t.Error("Its not working")
	} else {
		t.Log("Its working")
	}
}
func Test_StrAppend(t *testing.T){
	if StrAppend("reddy","sai")!="reddysai"{
		t.Error("Its not working")
	} else {
		t.Log("its working")
	}
}
func Test_StrTitle(t *testing.T){
	if StrTitle("reddy")!="Reddy"{
		t.Error("Its not working")
	} else {
		t.Log("its working")
	}
}
func Test_StrRepeat(t *testing.T){
	if StrRepeat("sai",2)!="saisai"{
		t.Error("Its not working")
	} else {
		t.Log("its working")
	}
}
func Test_StrReverse(t *testing.T){
	if StrReverse("sai")!="ias"{
		t.Error("Its not working")
	} else {
		t.Log("its working")
	}
}
func Test_StrSplit(t *testing.T){
	s:=StrSplit("reddy,sai")
	
	if s[0]!="reddy" || s[1]!="sai"{
		t.Error("its not working")
	} else{
		t.Log("Its working")
	}
}
func Test_StrSplitAfter(t *testing.T){
	s:=StrSplitAfter("reddy. sai")
	
	if s[0]!="reddy." || s[1]!=" sai"{
		t.Error("its not working")
	} else{
		t.Log("Its working")
	}
}
func Test_StrSplitN(t *testing.T){
	s:= StrSplitN("raja|rani",2)
	if s[0]!="raja" || s[1]!="rani" {
		t.Error("Its not working")
	} else{
		t.Log("Its working")
	}
}
func Test_StrSplitChar(t *testing.T){
	s:=StrSplitChar("sai")

	if s[0]!="s" || s[1]!="a" || s[2]!="i"{
		t.Error("Its not working")
	} else {
		t.Log("Its working")
	}
}
func Test_StrJoin(t *testing.T){
	s:=[]string{"reddy","sai"}
	if StrJoin(s)!="reddy sai"{
		t.Error("Its not working")
	} else{
		t.Log("Its working")
	}
}
func Test_StrFields(t *testing.T){
	s:=StrFields("how*,are you man!")
	t.Log(s)
	if s[0]!="how" || s[1]!="are" || s[2]!="you" || s[3]!="man"{
		t.Error("Its not working")
	} else{
		t.Log("Its working")
	}

}
//============
//following test cases are used by check.v1 package
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) Test_StrLength1(c *C) {
    c.Check(3, Equals,StrLength("sai"))
}
func (s *MySuite) Test_StrconvItoa1(c *C) {
    c.Check("3", Equals,StrconvItoa(3))
}
func (s *MySuite) Test_StrconvAtoi1(c *C){
	c.Check(3,Equals,StrconvAtoi("3"))
}
func (s *MySuite) Test_StrconvItof1(c *C){
	c.Check(3.0000,Equals,StrconvItof(3))
}
func (s *MySuite) Test_StrconvFtoi1(c *C){
	c.Check(3,Equals,StrconvFtoi(3.14))
}
func (s *MySuite) Test_StrAppend1(c *C){
	c.Check("reddysai",Equals,StrAppend("reddy","sai"))
}
func (s *MySuite) Test_StrTitle1(c *C){
	c.Check("Reddy Sai",Equals,StrTitle("reddy sai"))
}
func (s *MySuite) Test_StrRepeat1(c *C){
	c.Check("saisai",Equals,StrRepeat("sai",2))
}
func (s *MySuite) Test_StrReverse1(c *C){
	c.Check("ias",Equals,StrReverse("sai"))
}
func (s *MySuite) Test_StrSplit1(c *C){
	var s1=StrSplit("reddy,sai")
	c.Check(s1[0],Equals,"reddy")
	c.Check(s1[1],Equals,"sai")
}
func (s *MySuite) Test_StrSplitAfter1(c *C){
	var s1=StrSplitAfter("reddy. sai")
	c.Check(s1[0],Equals,"reddy.")
	c.Check(s1[1],Equals," sai")
}
func (s *MySuite) Test_StrSplitN1(c *C){
	var	s1=StrSplitN("raja|rani",2)
	c.Check(s1[0],Equals,"raja")
	c.Check(s1[1],Equals,"rani")
}
func (s *MySuite) Test_StrSplitChar1(c *C){
	var s1=StrSplitChar("sai")
	c.Check(s1[0],Equals,"s")
	c.Check(s1[1],Equals,"a")
	c.Check(s1[2],Equals,"i")
}
func (s *MySuite) Test_StrJoin1(c *C){
	s1:=[]string{"reddy","sai"}
	c.Check(StrJoin(s1),Equals,"reddy sai")
}
func (s *MySuite) Test_StrFields1(c *C){
	s1:=StrFields("how*,are you man!")
	c.Check(s1[0],Equals,"how")
	c.Check(s1[1],Equals,"are")
	c.Check(s1[2],Equals,"you")
	c.Check(s1[3],Equals,"ma")


}

