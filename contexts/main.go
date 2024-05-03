package main

import (
	"context"
	"fmt"
)

func contextTest(context context.Context) {
	fmt.Println("Context Test")
	fmt.Println(context)
	fmt.Println(context.Value("name"))

}
func main() {
	cxt := context.Background()
	cxt = context.WithValue(cxt, "name", "GO Contexts")
	contextTest(cxt)
}
