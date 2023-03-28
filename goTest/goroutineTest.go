package main

import (
	"fmt"
	"time"

)

var channel = make(chan string)

func sender() {
	channel <- "send Hello"
	fmt.Println("finish send")
	time.Sleep(time.Duration(2) * time.Second)
	channel <- "send Hello2"
	time.Sleep(time.Duration(1) * time.Second)
	close(channel)
}

func receive() {
	for receive := range channel {
		// receive := <-channel
		fmt.Printf("message:%v \n", receive)
	}
	panic("test")

}

func main() {
	go receive()
	go sender()
	time.Sleep(time.Duration(5) * time.Second)
	
	fmt.Printf("channel pointer : %v\n", &channel)
	fmt.Println("main finish")
}
