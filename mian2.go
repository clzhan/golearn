package main

import "fmt"

func main() {

	fmt.Println("hello world")

	event := make(chan struct{})
	c:= make(chan string)


	go func() {
		s:= <-c
		fmt.Println(s)
		close(event)
	}()

	c<- "hell"

	<-event

	fmt.Println("over")




}
