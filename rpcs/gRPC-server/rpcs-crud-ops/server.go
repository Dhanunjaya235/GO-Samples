package main

//protoc --go_out=../protogen/golang --go_opt=paths=source_relative orders/order.proto products/product.proto google/api/date.proto google/api/http.proto google/api/annotations.proto
// protoc --go_out=../protogen/golang --go_opt=paths=source_relative  --go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative  --grpc-gateway_out=../protogen/golang  --grpc-gateway_opt paths=source_relative,generate_unbound_methods=true  orders/order.proto products/product.proto google/api/date.proto google/api/http.proto google/api/annotations.proto
import (
	"fmt"
	"log"
	"net"
	"rpc/grpc/examples/protogen/golang/orders"
	"rpc/grpc/examples/service"

	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", ":8883")

	if err != nil {
		log.Fatalf("failed to listen %v", err.Error())
	}

	server := grpc.NewServer()

	db := service.NewDB()
	orderService := service.NewOrderService(db)

	orders.RegisterOrdersServer(server, &orderService)
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		fmt.Printf("failed to serve %v", err.Error())
	}
}
