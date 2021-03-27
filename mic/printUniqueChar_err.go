package main

import "fmt"
//malayalam
func PrintUnique(in string ){
	charMap:=make(map[rune]int)
	//string to runes
	for i:=0; i <len(in);i++{
		charMap[in[i]]++
	}

	for k,v:= range charMap{
		if v==1{
			fmt.Println(k)
			break
		}
	}

}