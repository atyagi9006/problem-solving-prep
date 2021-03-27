package main

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

func main_1() {
	email := "5aas.qe+u_kjh_j009k@gmail.com"
	fmt.Println("result",govalidator.IsEmail(email))
	if !govalidator.IsEmail(email) {
		fmt.Println("Not a email")
	}
}

type my struct{
	name string
}

type my2 struct{
	
}
func main(){
	d:=make(chan bool)
	strchan:=make(chan string)

	fmt.Println("Starting work..")
	go doSomeThing(d, strchan)
	strchan<-"hello"
	fmt.Println("Working on something else")

	
	//blocking till the works finishes
/* 	for v :=range d{
		fmt.Println(v)

	} */
	<-d
	
	close(d)
}

func doSomeThing(done chan bool, strchan chan string){
	time.Sleep(3*time.Second)
	fmt.Println(<-strchan)
	fmt.Println("work done ")
	done <- true
}


