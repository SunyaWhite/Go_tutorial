package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("Enter some number : ")
	fmt.Scan(&input)
	num, err := strconv.Atoi(input)
	if err != nil {
		panic("Error occured : " + err.Error())
	}
	fmt.Println(num + 1)
}
