package main

import (
	"fmt"
	"net/rpc"
)

type Numbers struct {
	Num1, Num2 int
}

type Result int

func main() {
	client, err := rpc.Dial("tcp", "localhost:2345")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer client.Close()

	numbers := Numbers{Num1: 10, Num2: 20}

	var (
		sum     Result
		product int
	)

	additionError := client.Call("Calculator.Addition", numbers, &sum)
	if additionError != nil {
		fmt.Println(additionError.Error(), "Addition Error")
	}

	multiplyError := client.Call("Calculator.Multiplication", numbers, &product)
	if multiplyError != nil {
		fmt.Println(multiplyError.Error(), "Multiplication Error")
	}

	fmt.Printf("Sum Of %d and %d = %d \n", numbers.Num1, numbers.Num2, sum)
	fmt.Printf("Product Of %d and %d = %d \n", numbers.Num1, numbers.Num2, product)

}
