package main
/*
a=[a,b,d]
b=[b,c,a,b]

c=[c]
c= (a+b)
*/
func unique(a,b [] string)[]string{
 res :=make([]string,0)
 amap:=make(map[string]int)
	for _,v:=range a{ //a len m
		amap[v]=1
	}

	for _,v:= range b{//b lne n
		if _,ok:=amap[v]; !ok{
			res= append(res, v)
		}
	}

//complexicity m+n
 return res
}