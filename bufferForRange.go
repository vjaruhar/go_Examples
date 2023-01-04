package main

import "fmt"

func foo(c chan int) {
	fmt.Println("ok")
	num := <-c
	fmt.Println(num)
	//close(c)
}

func foo1(c chan int) {
	fmt.Println("dusra routine")
	num := <-c
	fmt.Println(num)
}

func main() {

	c := make(chan int)

	//go foo(c)
	go foo1(c)
	go foo(c)
	go foo1(c)
	go foo1(c)

	c <- 1
	c <- 2
	c <- 3

	c <- 4
	close(c)

	for val := range c {
		fmt.Println(val)
	}
}
