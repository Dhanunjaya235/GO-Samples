package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Calculator struct{}

type Numbers struct {
	Num1, Num2 int
}

type Result int

func (c *Calculator) Addition(nums *Numbers, result *Result) error {

	*result = Result(nums.Num1) + Result(nums.Num2)
	return nil
}

func (c *Calculator) Multiplication(nums *Numbers, result *int) error {

	*result = nums.Num1 * nums.Num2
	return nil
}

func main() {

	calculator := new(Calculator)

	rpc.Register(calculator)

	listener, err := net.Listen("tcp", ":2345")

	if err != nil {
		fmt.Println(err.Error())
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err.Error())
		}

		go rpc.ServeConn(conn)

	}

}
