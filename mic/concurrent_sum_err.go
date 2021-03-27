package main


func process(input []int){
	l:=len(input)
	res:=make(chan int,3)
	 if l>1000{
		 sum(input,res)
	 }
	 
	in1:=input[:l/4]
	in2:=input[(l/4)+1:l/2]
	in3:=input[(l/2)+1:]

	 go sum(in1,res)
	 go sum(in2, res)
	 go sum(in3,res)

	 total:=0
	 for s:=range res{
		total+=s
	 }


}


func sum(input[]int, res chan int) int{
	sum:=0
	for i:=0;i<len(input);i++{
		sum+=input[i]
	}
	res<- sum
	return sum
}