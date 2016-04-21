package stringutil

import (
    "testing" 
    )
 
func Test_StrSuffix(t *testing.T) {
 
	if StrSuffix("aruna","a")==true {
		t.Log("Test passed")
	}else {
		t.Error("StrSuffix did not work as expected.")
	}
}

func Test_StrPrefix(t *testing.T) {

	if StrPrefix("aruna","ar")==true {
		t.Log("Test passed")
	}else {
		t.Error("StrPrefix did not work as expected.")
	}
}


func Test_StrCompareLen(t *testing.T) {

	if StrCompareLen("aruna","aruna")==true {
		t.Log("Test passed")
	}else {
		t.Error("StrCompareLen did not work as expected.")
	}
}

func Test_StrCompare(t *testing.T) {

	if StrCompareLen("aruna","aruna")==true {
		t.Log("Test passed")
	}else {
		t.Error("StrCompare did not work as expected.")
	}
}

func Test_StrContains(t *testing.T) {

	if StrContains("aruna","aruna")==true {
		t.Log("Test passed")
	}else {
		t.Error("StrContains did not work as expected.")
	}
}

func Test_StrUpper(t *testing.T) {

	if StrUpper("aruna")=="ARUNA" {
		t.Log("Test passed")
	}else {
		t.Error("StrUpper did not work as expected.")
	}
}

func Test_StrLower(t *testing.T) {

	if StrLower("aruna")=="aruna" {
		t.Log("Test passed")
	}else {
		t.Error("StrLower did not work as expected.")
	}
}

func Test_StrReplace(t *testing.T) {

	if StrReplace("aruna how are you",1)=="aruna,how are you" {
		t.Log("Test passed")
	}else {
		t.Error("StrReplace did not work as expected.")
	}
}

func Test_StrSwap(t *testing.T) {
	one,sec:= StrSwap("arun","kumar")
	if one=="kumar" && sec=="arun"{
		t.Log("Test Passed")
	}else {
		t.Error("StrSwap did not work as expected.")
	}
}

func Test_StrSort(t *testing.T) {
   s:=[]string{"cat","ate","dog"}
    s1:=StrSort(s)
   if s1[0]=="ate" && s1[1]=="cat" &&s1[2]=="dog" {
   	t.Log("Test Passed")
   }else {
   	t.Error("StrSort did not work as expected")
   }
}

func Test_StrFirstWords(t *testing.T) {

   //value := "there are many reasons"
   if StrFirstWords("there are many reasons",1)=="there" {
   		t.Log("Test Passed")
   }else {
   	t.Error("StrFirstWords did not work as expected")
   }
}

func Test_StromitBack(t *testing.T) {

   if StromitBack(1,"arun")=="a" {
   	t.Log("Test Passed")
   }else {
   	t.Error("StromitBack did not work as expected")
   }
}

func Test_StromitFront(t *testing.T) {

   if StromitFront(1,"arun")=="run" {
   	t.Log("Test Passed")
   }else {
   	t.Error("StromitFront did not work as expected")
   }
}

func Test_StrSubstr(t *testing.T) {

   if StrSubstr(5,"arun:kumar")=="kumar" {
   	t.Log("Test Passed")
   }else {
   	t.Error("StrSubstr did not work as expected")
   }
}

func Test_StrReadLines(t *testing.T) {
	s:=StrReadLines("arun.txt")
	t.Log(s)
  if s[0]=="hello"{
    t.Log("Test Passed")
   }else {
   	t.Error("StrSubstr did not work as expected")
   }
}