package main

import (
	"fmt"
	"math/rand"
	"time"
)

// function to write
func merge2ChannelsFunc(fun func(int) int, a <-chan int, b <-chan int, out chan<- int, n int){
	go func(){

		bufA := make([] int, 0)
		bufB := make([] int, 0)

		countA := 0
		countB := 0
		countOut := 0
		//defer close(out)

		for countOut != n{
			select {
				case v, ok := <- a:
					if !ok || countA == n{
						a = nil
						continue
					}
					countA++
					bufA = append(bufA, v)
					if countA == n{
						a = nil
						continue
					}
				case v, ok := <- b:
					if !ok || countB == n{
						b = nil
						continue
					}
					countB++
					bufB = append(bufB, v)
					if countB == n{
						b = nil
						continue
					}
				default:
					if len(bufA) != 0 && len(bufB) != 0{
						fmt.Println("Buffer A : ", bufA)
						fmt.Println("Buffer B : ", bufB)
						/*go func(a int, b int) {
							innerOut := make(chan  int)
							go executeSomeHeavyOperation(fun, innerOut, a)
							go executeSomeHeavyOperation(fun, innerOut, b)
							res1 := <-innerOut
							res2 := <-innerOut
							out <- res1 + res2
							close(innerOut)
							innerOut = nil // is it necessary?
							countOut++
						}(bufA[0], bufB[0]) The way I would do it */
						innerOut := make(chan  int)
						go executeSomeHeavyOperation(fun, innerOut, bufA[0])
						go executeSomeHeavyOperation(fun, innerOut, bufB[0])
						res1 := <-innerOut
						res2 := <-innerOut
						out <- res1 + res2
						close(innerOut)
						innerOut = nil // is it necessary?
						countOut++
						bufA = bufA[1:]
						bufB = bufB[1:]
					}
					if countOut == n{
						close(out)
					}

			}
		}
	}()

	fmt.Println("Merge function is set")
	// start out function
	//go mergerFunc()
	fmt.Println("Merge function is started")
}

func executeSomeHeavyOperation(fun func(int) int, a chan<- int, arg int){
	a <- fun(arg)
}

func someHeavyOperation(a int) int{
	fmt.Println("Waiting for some operation to complete with num: ", a)
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	return a*a
}

func randomNumberGenerator(c chan<- int, n int){
	time.Sleep(time.Second)
	fmt.Println("Generator is started")
	for i := 0; i< n; i++{
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		num := rand.Intn(5)
		fmt.Println("Sending new number into channel : ", num)
		c <- num

	}
	close(c)
	fmt.Println("Generator is done")
}

func main() {
	n := 6

	countI := 0

	in1 := make(chan int, n)
	in2 := make(chan int, n)
	out := make(chan int, n)

	go randomNumberGenerator(in1, n + 4)
	go randomNumberGenerator(in2, n + 2)

	merge2ChannelsFunc(someHeavyOperation, in1, in2, out, n)

	for out != nil {
		select{
			case v, ok := <-out:
				if !ok{
					out = nil
					break
				}
				fmt.Println("The result is ", v)
				countI++
			default:
				countI++
				time.Sleep(time.Second)
				fmt.Println("The main thread is working. Passed ", countI, " seconds" )
		}
	}

	time.Sleep(time.Second)

}
