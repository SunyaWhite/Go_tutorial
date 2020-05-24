package main

import (
	"fmt"
	"time"
)

var someGlobalFunc func()

func printHello(){
	fmt.Println("Hello, motherfucking world")
}

func funcAssigment(){
	someGlobalFunc = func(){
		time.Sleep(time.Second)
		fmt.Println("Hello, motherfucking world")
	}
	go someGlobalFunc()
	fmt.Println("Hello from funcAssigment")
}

func main() {
	funcAssigment()
	time.Sleep(time.Second * 3)
	fmt.Println("Done!")
}
