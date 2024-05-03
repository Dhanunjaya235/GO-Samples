package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(number int) bool {
	return number%2 == 0
}

func withOutMutex() {
	n := 0

	go func() {
		isEven := isEven(n)
		time.Sleep(1 * time.Second)
		if isEven {
			fmt.Printf("With Out Mutex Value of n is %d and it is even \n", n)
			return
		}
		fmt.Printf("With Out Mutex Value of n is %d and it is odd \n", n)
	}()

	go func() {
		n++
	}()
}

func withMutex() {
	n := 0
	var mutex sync.Mutex

	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		isEven := isEven(n)
		time.Sleep(1 * time.Second)
		if isEven {
			fmt.Printf("With Mutex Value of n is %d and it is even \n", n)
			return
		}
		fmt.Printf("With Mutex Value of n is %d and it is odd \n", n)
	}()

	go func() {
		mutex.Lock()
		n++
		fmt.Println("GO Routine 2 in withMutex")
		mutex.Unlock()
	}()
}

func main() {
	withOutMutex()
	withMutex()
	time.Sleep(3 * time.Second)
}
