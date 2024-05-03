package main

import (
	"fmt"
	"math"

	"sample-mod/calculator"

	"encoding/json"

	"rsc.io/quote/v4"
	"rsc.io/sampler"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {

	message := make(chan any)
	message2 := make(chan any)
	go func() {
		fmt.Println("Go Routine2")
		message2 <- "ping 2"
	}()
	go func() {
		fmt.Println("Go Routine1")
		message <- "ping 1"
	}()

	fmt.Println("Calling Function")
	msg := <-message
	msg2 := <-message2
	fmt.Println(msg)
	fmt.Println("Calling Function2")
	fmt.Println(msg2)

	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
	fmt.Println(quote.Hello())

	test()
	deferTest()
	fmt.Println(sampler.Hello())

	jsonData()

}

func test() {
	sum, err := calculator.Addition(10, 20)
	if err != nil {
		fmt.Println((err.Error()))
	} else {
		fmt.Println(sum)
	}

	product, error := calculator.Multiply(10)
	if error != nil {
		fmt.Println((error.Error()))
	} else {
		fmt.Println(product)
	}

	num1, num2 := 10, 10
	quotient, divError := calculator.Division(num1, num2)

	if divError == nil {
		fmt.Printf("The division of %d / %d is %d", num1, num2, quotient)
	} else {
		fmt.Println(divError.Error())
	}
}

func deferTest() {
	fmt.Println("Before Defer Line")

	defer fmt.Println("Defer Line")

	fmt.Println("After Defer Line")
}

func jsonData() {
	m := Message{"Alice", "Hello", 1294706395881547000}

	b, err := json.Marshal(m)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(b))
	var data map[string]interface{}

	if err := json.Unmarshal(b, &data); err != nil {
		fmt.Println(err.Error())
		return
	}

	for key, value := range data {
		fmt.Printf("Key: %s \t Value: %v   \n", key, value)
	}

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
