package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"rpc/grpc/examples/protogen/golang/orders"
)

func main() {
	conn, err := grpc.Dial("localhost:8883", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
	}

	defer conn.Close()

	mux := runtime.NewServeMux()

	if err = orders.RegisterOrdersHandler(context.Background(), mux, conn); err != nil {
		fmt.Println(err.Error())
	}

	addr := "localhost:6969"
	fmt.Println("API gateway server is running on " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		fmt.Println(err.Error())
	}

}
