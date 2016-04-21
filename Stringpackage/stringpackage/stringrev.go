package main 

import ("fmt"
        "./string"
        //"sort"
        )

func main() {
	 
	c:=stringutil.StrReverse("abcdefghijklmanopqrstuvwxyz")
    fmt.Println("StrReverse - ",c)
   
    d:=stringutil.StrLength("c")
    fmt.Println("StrLength - ",d)
   
    e:=stringutil.StrconvItoa(42)
    fmt.Printf("%T",e)
    fmt.Println("StrconvItoa - ",e)
 
    f:=stringutil.StrconvAtoi("42")
    fmt.Printf("%T",f)
    fmt.Println("StrconvAtoi - ",f)
   
    g:=stringutil.StrconvItof(13)
    fmt.Printf("%T",g)
    fmt.Println("StrconvItof - ",g)
   
    h:=stringutil.StrconvFtoi(13.132)
    fmt.Printf("%T",h)
    fmt.Println("StrconvFtoi - ",h)
    
    i:=stringutil.StrAppend("arun","kumar")
    fmt.Println("StrAppend - ",i)
   
    j:=stringutil.StrTitle("the arun")
    fmt.Println("StrTitle - ",j)
    
    k:=stringutil.StrRepeat("abc",3)
    fmt.Println("StrRepeat - ",k)
    
    l:=stringutil.StrSplit("cat,dog,bat")
    fmt.Println("StrSplit - ",l)

    p:=stringutil.StrSplitAfter("cat.dog.bat")
    fmt.Println("StrSplitAfter - ",p)

    m:=[]string{"cat","dog","bat"}
    fmt.Println(stringutil.StrJoin(m))

    n:=stringutil.StrFields("arun kumar   /*-+/ aaron,cranky")
    fmt.Println("StrFields - ",n)
     
    o:=stringutil.StrReadLines("arun.txt")
    fmt.Println("StrReadLines - ",o)
    
    q:=stringutil.StrSplitN("12|13|14|15|16",5)
    fmt.Println("StrSplitN - ",q)

    r:=stringutil.StrSplitChar("hello")
    fmt.Println("StrSplitChar - ",r)

    t:=stringutil.StrSubstr(4,"cat;dog")
    fmt.Println("StrSubstr - ",t)

    u:=stringutil.StromitFront(1,"arunkumar")
    fmt.Println("StromitFront - ",u)

    v:=stringutil.StromitBack(2,"arun")
    fmt.Println("StromitBack - ",v)

    value := "there are many reasons"

    result1 := stringutil.StrFirstWords(value, 1)
    fmt.Println(result1)

    result2 := stringutil.StrFirstWords(value, 2)
    fmt.Println(result2)

    result3 := stringutil.StrFirstWords(value, 3)
    fmt.Println(result3)
     
    result4 := stringutil.StrFirstWords(value, 10)
    fmt.Println(result4)

    animals:=[]string{"dog","cat","ate"}
    w:=stringutil.StrSort(animals)
    fmt.Println("StrSort - ",w)

    a,b:=stringutil.StrSwap("kumar","arun")
    fmt.Println("StrSwap - ",a,b)

    replace:=stringutil.StrReplace("Hey Dude-What's up Mann Having Fun Ha ",1)
    fmt.Println("StrReplace - ",replace)

    strupper:=stringutil.StrUpper("Hey Dude-What's up")
    fmt.Println("StrUpper - ",strupper)

    strlower:=stringutil.StrLower("Hey Dude-What's up")
    fmt.Println("StrLower - ",strlower)
    
    strcontains:=stringutil.StrContains("whats up dude ", "up")
    fmt.Println("StrContains - ",strcontains)

    strcompare:= stringutil.StrCompare("arun","arun")
    fmt.Println("StrCompare - ",strcompare)
     
    strcomparelen:= stringutil.StrCompareLen("aruna","arun")
    fmt.Println("StrCompareLen - ",strcomparelen)

    strprefix:= stringutil.StrPrefix("aruna","ar")
    fmt.Println("StrPrefix - ",strprefix)

    strsuffix:= stringutil.StrSuffix("aruna","a")
    fmt.Println("StrSuffix - ",strsuffix)

}