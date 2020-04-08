// generator_fib4.go
//
// use a channel to simulate a generator
//
// channels are the pipes that connect concurrent goroutines
// send to channel syntax is "chan <- v"
// receive from channel syntax is "v, status := <-chan"
// (data flows in the direction of the arrow)
// receive status is true unless closed with closed(chan)
//
// more detailed info at:
// http://golang.org/ref/spec#Making_slices_maps_and_channels
// https://golang.org/ref/spec#Receive_operator
// online play at:
// http://play.golang.org/p/8CrOj6lmKe
//
// tested with Go version 1.4.2   by vegaseat  3may2015
package main

import "fmt"

// generator using a channel and a goroutine
func fib(n int) chan int {
	// create a channel
	c := make(chan int)
	// create an anonymous inner function
	// keyword "go" starts a goroutine
	go func() {
		x, y := 0, 1
		for k := 0; k < n; k++ {
			// send value x to the channel c
			c <- x
			// swap
			x, y = y, x+y
		}
		// close(c) sets the status of the channel c to false
		// and is needed by the for/range loop to end
		close(c)
	}()
	return c
}
func main() {
	fmt.Println("A Fibonacci generator via channel/goroutine:")
	// function fib() returns a channel
	// so fc would be of type "chan int"
	// sum := 0
	fc := fib(20)
	fmt.Printf("fc is type %T\n", fc)
	k := 0
	n:=0
	for k = range fc {
		
			fmt.Println(n, k)

		n++

	}
	fmt.Println("Index is", k)
}
