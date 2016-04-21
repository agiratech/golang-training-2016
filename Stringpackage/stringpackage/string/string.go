package stringutil

import ("strconv"
        "strings"
        "bufio"
        "sort"
        "os"
        //"regexp"
      // "fmt"
       
        )

func StrLength(s string) int {

	c:=len(s)
	return c 
}

func StrconvItoa(s int) string {
   
   a:=strconv.Itoa(int(s))
   return a 

}

func StrconvAtoi(s string) int {
   
   a,_:=strconv.Atoi(string(s))
   return a 

}

func StrconvItof(s int) float64 {

	a:=float64(s)
	return a 
}


func StrconvFtoi(s float64) int {

	a:=int(s)
	return a 
}

func StrAppend(a ,s string) string {

	b:=a + s
	return b
}

func StrTitle(s string)string {

	a:= strings.Title(s)
	return a 
}

func StrRepeat(s string,n int)string {
    
    var a string
    for i:=0; i<=n; i++ {
	a=strings.Repeat(s, i)
    }
	return a 
}

func StrReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func StrSplit(s string) []string {
    
	a := strings.Split(s,",")
    return a
}

func StrSplitAfter(s string)[]string{

	a:= strings.SplitAfter(s,".")
	return a 

}

func StrSplitN(s string,i int) []string {
   
   a:=strings.SplitN(s,"|",i)
   return a
   
}

func StrSplitChar(s string) []string{

   a:=strings.SplitAfter(s,"")
   return a 

}

func StrJoin(s []string) string {

	a:= strings.Join(s," ")
	return a

}

func StrFields(s string) []string {

	 f := func(c rune) bool {
	return c == ' ' || c == ',' || c == ':' || c =='.' || c == '|' || c == '+' || c == '-' || c == '*' || c == '/' || c == '!' || c == '$' || c == '%' || c == '&' || c == '^' || c == '#' || c == '@' || c == '_'
    }
     fields := strings.FieldsFunc(s, f)
     return fields
}

func StrReadLines(s string) []string {
  
  a,_:= os.Open(s)
  var lines []string
  scanner := bufio.NewScanner(a)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines
}

func StrSubstr(i int, s string) string {

	substring:= s[i:len(s)]
	return substring
}

func StromitFront(i int,s string) string {

    substring:= s[i:]
    return substring
}

func StromitBack(i int,s string) string {

	substring:= s[:i]
	return substring
}
 
func StrFirstWords(s string, i int) string {

    for k:= 0 ; k < len(s) ; k++ {
    	if s[k] == ' ' {
    		i = i - 1

    		if i == 0 {
    			return s[0:k]
    		}
    	}
    } 
return s
}

func StrSort(s []string) []string {
     
    sort.Strings(s)
    return s 

}

func StrSwap(x , y string) (string, string) {
	
	var temp string

	temp = x 
	x = y 
	y = temp

	return x, y ;

}

func StrReplace(s string, i int) string {

	str := strings.Replace(s, " ", ",", i)
	return str
}

func StrLower(s string) string {

    str := strings.ToLower(s)
    return str

}

func StrUpper(s string) string {

    str := strings.ToUpper(s)
    return str

}

func StrContains(s ,t string) bool {

   str:=strings.Contains(s, t)
   return str

}

func StrCompare(s1 string, s2 string) bool {
      
  var v bool=false

  if s1 == s2 {
  v = true
    }
 return v
} 

func StrCompareLen(s1 string, s2 string) bool {
      
  var v bool=false

  if len(s1) == len(s2) {
  v = true
  }
return v
} 

func StrPrefix(s, t string) bool {
 
 str:=strings.HasPrefix(s,t)
 return str

}

func StrSuffix(s, t string) bool {
 
 str:=strings.HasSuffix(s,t)
 return str

}