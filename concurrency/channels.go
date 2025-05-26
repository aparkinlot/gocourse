package concurrency

import "fmt"

// Channels cannot be compiled/used without a goroutine feeding the channel
// -> fatal error: all goroutines are asleep - dealock!

// Channels are blocking whereas Goroutines are non-blocking
// Concurrency errors are found at runtime (go runtime scheduler works with these concepts)

func test_channels() {
	// var := make(chan <type>)
	greeting := make(chan string) // unbuffered channel, can take many values into this and reciever takes them out FIFO
	greet := "Hello"

	go func() { // wihtout the go func, the main thread would not be able to continue
		greeting <- greet // blocking because it is continuously trying to recieve values, it is ready to receive continuous flow of data
	}()

	receiver := <-greeting // new var recieving the same input/pipe value as greet var
	// if no value is received, will block until a value is received
	// this is part of the main goroutine
	// for any further values send to the channel, a subsequent receiver must be established
	fmt.Println(receiver, greet)
}
