package main 

import ("fmt"
        "./string"
        //"sort"
        )

func main() {
	 
	c:=stringutil.StrReverse("abcdefghijklmanopqrstuvwxyz")
    fmt.Println(c)
   
    d:=stringutil.StrLength("c")
    fmt.Println(d)
   
    e:=stringutil.StrconvItoa(42)
    fmt.Printf("%T",e)
    fmt.Println(e)
 
    f:=stringutil.StrconvAtoi("42")
    fmt.Printf("%T",f)
    fmt.Println(f)
   
    g:=stringutil.StrconvItof(13)
    fmt.Printf("%T",g)
    fmt.Println(g)
   
    h:=stringutil.StrconvFtoi(13.132)
    fmt.Printf("%T",h)
    fmt.Println(h)
    
    i:=stringutil.StrAppend("arun","kumar")
    fmt.Println(i)
   
    j:=stringutil.StrTitle("the arun")
    fmt.Println(j)
    
    k:=stringutil.StrRepeat("abc",3)
    fmt.Println(k)
    
    l:=stringutil.StrSplit("cat,dog,bat")
    fmt.Println(l)

    p:=stringutil.StrSplitAfter("cat.dog.bat")
    fmt.Println(p)

    m:=[]string{"cat","dog","bat"}
    fmt.Println(stringutil.StrJoin(m))

    n:=stringutil.StrFields("arun kumar   /*-+/ aaron,cranky")
    fmt.Println(n)
     
    o:=stringutil.StrReadLines("arun.txt")
    fmt.Println(o)
    
    q:=stringutil.StrSplitN("12|13|14|15|16",5)
    fmt.Println(q)

    r:=stringutil.StrSplitChar("hello")
    fmt.Println(r)

    t:=stringutil.StrSubstr(4,"cat;dog")
    fmt.Println(t)

    u:=stringutil.StromitFront(1,"arunkumar")
    fmt.Println(u)

    v:=stringutil.StromitBack(2,"arun")
    fmt.Println(v)

    value := "there are many reasons"

    result1 := stringutil.StrFirstWords(value, 1)
    fmt.Println(result1)

    animals:=[]string{"dog","cat","ate"}
    w:=stringutil.StrSort(animals)
    fmt.Println(w)

    a,b:=stringutil.StrSwap("kumar","arun")
    fmt.Println(a,b)

    replace:=stringutil.StrReplace("Hey Dude-What's up Mann Having Fun Ha ",1)
    fmt.Println(replace)

    strupper:=stringutil.StrUpper("Hey Dude-What's up")
    fmt.Println(strupper)

    strlower:=stringutil.StrLower("Hey Dude-What's up")
    fmt.Println(strlower)
    
    strcontains:=stringutil.StrContains("whats up dude ", "up")
    fmt.Println(strcontains)

    strcompare:= stringutil.StrCompare("arun","arun")
    fmt.Println(strcompare)
     
    strcomparelen:= stringutil.StrCompareLen("aruna","arun")
    fmt.Println(strcomparelen)

    strprefix:= stringutil.StrPrefix("aruna","ar")
    fmt.Println(strprefix)

    strsuffix:= stringutil.StrSuffix("aruna","a")
    fmt.Println(strsuffix)

}