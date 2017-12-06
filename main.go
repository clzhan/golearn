package main

import (
	"fmt"
	"time"
)

//const LengthW ,LengthH  = 1, "123"

const(
	LengthW int = 1
	LengthH string="123"
)

const(
	x=iota
	y
	z
)

func test(a int, b int) int {
	return  a+b

}

func  Task1()  {
	fmt.Println("Task1 start")
	time.Sleep(1* time.Second)
	fmt.Println("Task1 end")

}
func  Task2()  {
	fmt.Println("Task2 start")
	time.Sleep(2* time.Second)
	fmt.Println("Task2 end")

}

func main()  {
	x := [4] int{0,1,2,3}


	fmt.Println(time.Now())

	go Task1()
	Task2()

	go func() {
		fmt.Println("Task3 start")
		time.Sleep(2* time.Second)
		fmt.Println("Task3 end")
	}()

	fmt.Println(time.Now())

	fmt.Println(x)


}
