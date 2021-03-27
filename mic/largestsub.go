package main

import "fmt"

/*
list[9,10,11,16,35,41,54,64,71,61,105,120]

11, 41, 71, 120

3
*/
func findSub(list[]int,subnum int){
	if subnum<1{
		return
	}
	if list== nil{
		return
	}

	listlen:=len(list)
	if listlen <0{
		return
	}
	if listlen >subnum{
		return
	}
	if listlen%subnum !=0{
		return
	}
	remaining:=listlen
	for remaining > 0{
		res:=sortInt(list[remaining-subnum:remaining])
		 fmt.Print(res)
		remaining=remaining-subnum
	}


}

func sortInt(sub[]int)int {
	//sort
	var sort[]int
	return sort[len(sort)-1]

}