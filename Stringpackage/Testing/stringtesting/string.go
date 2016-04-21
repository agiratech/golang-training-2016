package stringutil

import ("strings"
       "sort"
       "os"
       "bufio"
       )


func StrSuffix(s, t string) bool {
 
 str:=strings.HasSuffix(s,t)
 return str

}

func StrPrefix(s, t string) bool {
 
 str:=strings.HasPrefix(s,t)
 return str

}

func StrCompareLen(s1 string, s2 string) bool {
      
  var v bool=false
  if len(s1) == len(s2) {
  v = true
  }
return v
} 

func StrCompare(s1 string, s2 string) bool {
      
  var v bool=false

  if s1 == s2 {
  v = true
    }
 return v
} 

func StrContains(s ,t string) bool {

   str:=strings.Contains(s, t)
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

func StrReplace(s string, i int) string {

	str := strings.Replace(s, " ", ",", i)
	return str
}

func StrSwap(x , y string) (string, string) {
	
	var temp string

	temp = x 
	x = y 
	y = temp

	return x, y 
}

func StrSort(s []string) []string {
     
    sort.Strings(s)
    return s 

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

func StromitBack(i int,s string) string {

	substring:= s[:i]
	return substring
}

func StromitFront(i int,s string) string {

    substring:= s[i:]
    return substring
}

func StrSubstr(i int, s string) string {

	substring:= s[i:len(s)]
	return substring
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