package main

import (
    "fmt"
    "strings"
)

func main() {
	str := "red,?"
	fmt.Println("cou---",strings.ContainsAny(str, ","))
	if strings.ContainsAny(str, ","){
		str = strings.Replace(str, ",", "", -1)
	}
	if strings.ContainsAny(str, "?"){
		str = strings.Replace(str, "?", "", -1)
	}
   // res=strings.Replace(str,".","",1)
    //res=strings.Replace(str,"-","",1)
    //res=strings.Replace(str,"?","",1)
    fmt.Println(str)
}