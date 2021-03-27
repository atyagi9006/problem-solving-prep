package main

func main(){

}

type advice struct{
	buyAt int
	sellAt int
}

func stockAdvice(prices []int){
	lenPr:=len(prices)
	min, minINdex:=findMin(prices)





	
}

func findMin(input[]int )(int,int){
	min:=0
	index:=0
	if input[0]>input[1]{
		min=input[1]
		index=1
	}else{
		min=input[0]
		index=0
	}
	for i:=2;i<len(input);i++{
		if input[i]<min {
			min=input[i]
			index=i
		}
	}
	return min,index
}
