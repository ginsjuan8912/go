package main

import "fmt"

const num = 3

// Producer /*The following function is a producer that will generate a number a put it on a write-only channel*/
func Producer(channel chan<- int) {
	for i := 0; i < num; i++ {
		channel <- i
		fmt.Println(i, "Sent")
	}
}

func Consumer(channel <-chan int) {
	for i := 0; i < num; i++ {
		number := <-channel
		fmt.Println("Recieved", number)
	}
}

func main() {
	//The following syntax builds a channel
	channel := make(chan int)

	go Producer(channel)
	Consumer(channel)

}