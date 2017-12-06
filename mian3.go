package main

import "fmt"

func main() {
	c:=make(chan int,3)
	c<- 1
	c<-2

	fmt.Println(<-c)
	c<-3
	c<-4
	fmt.Println(<-c)
}
