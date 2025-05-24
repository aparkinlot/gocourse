package main

import (
	"fmt"
	"time"
)

// Goroutines are just funcitons that leave the main thread and run in the background and
// come back to join the main thread once the functions are finshed/ready to return any values

// Goroutines do not stop the program flow and are non-blocking
// separate and independent functions/threads

func test_goroutines() { // main always runs in the main thread
	var err error
	fmt.Println("Begin")
	go hello()                                                                     // runs on a separate thread -> main thread stops and this func/thread runs immediately
	fmt.Println("hello function has run and main thread continues (non-blocking)") // proceed to this

	go printNums()
	go printLetters()

	// these two threads will run when scheduled -> out of user control and therefore, may have varying results that are printed

	go func() {
		err = work()
	}()

	time.Sleep(1 * time.Second) // if this is commented out -> goroutine leak as we do not get the error

	if err == nil {
		fmt.Println("main thread finished before work thread could return error")
	} else {
		fmt.Println("work thread finished, here is the error:", err)
	}
}

func hello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNums() {
	for i := range 5 {
		fmt.Println("Number:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("Char:", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func work() error {
	// simulate computation

	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occured in work func")
}
