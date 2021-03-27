package main

import "fmt"

func main() {
	work:=make(chan int,100)
	result:=make(chan int, 100)

	//starting few workers
	for s:=1;s<4;s++{
		go subscriberDoSq(s,work,result)
	}

	//call publisher
	 publisher(work)
	
	 close(work)


	 for i:=1;i<=100;i++{
		 fmt.Println("Result: ",<-result)
	 }


}


func publisher(work chan int){
	for i:=1;i<=100;i++{
		work <-i
	}
}
func subscriberDoSq(id int, work <-chan int, result chan<- int) {
	for w := range work {
		fmt.Printf("Subscriber %d working on %v \n", id, w)
		result <- w * w
		fmt.Printf("Subscriber %d finished on %v \n",id,w)
	}
}
