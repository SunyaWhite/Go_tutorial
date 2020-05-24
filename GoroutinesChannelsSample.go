package main

import (
	"fmt"
	"math/rand"
	"time"
)

func someCalculations(index int, in chan int) int {
	//sec := rand.Int() % 10
	sec := <- in
	fmt.Println("Wait for goroutine ", index, " sec ", sec)
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println("Done for goroutine ", index)
	return rand.Int() % 5
}

func main() {
	var in chan int = make(chan int)
	fmt.Println("Hello ")
	go someCalculations(1, in)
	go someCalculations(2, in)
	time.Sleep(time.Second * 6)
	in <- 3
	in <- 4
	time.Sleep(time.Second * 5)
	fmt.Println("Done")
}
