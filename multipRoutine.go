package main

import "fmt"

func square(c chan int) {
	fmt.Println("Square reading..")
	//num := <-c
	c <- 6
}

func cube(c chan int) {
	fmt.Println("Cube reading..")
	num := <-c
	c <- num * num * num
}
func lena(c chan int) {
	//c <- 3
}

func main() {
	fmt.Println("main() started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)
	go lena(cubeChan)

	testNum := 3
	fmt.Println("using test number 3, routines started")
	fmt.Println("Square routine started")

	squareChan <- testNum
	//cubeChan <- 3
	fmt.Println("------")
	//cubeChan <- 3

	for {
		val := <-cubeChan
		if val == 3 {
			fmt.Println("inside for")
			break
		}

	}
	fmt.Println("Main resumed, starting cube")
	/*
		cubeChan <- 3
		fmt.Println("Resuming main..")

		squareVal, cubeVal := <-squareChan, <-cubeChan

		sum := squareVal + cubeVal
		fmt.Println("Sum of square and cube of %v is %v", testNum, sum)
	*/

}
