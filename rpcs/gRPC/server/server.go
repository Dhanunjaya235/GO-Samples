package main

//protoc --go_out=../protogen/golang --go_opt=paths=source_relative orders/order.proto products/product.proto google/api/date.proto

import (
	"encoding/json"
	"fmt"
	"log"
	"rpc/grpc/examples/database"
	"rpc/grpc/examples/protogen/golang/orders"
	product "rpc/grpc/examples/protogen/golang/products"

	"google.golang.org/genproto/googleapis/type/date"
)

func main() {

	orderItem := orders.Order{
		OrderId:    10,
		CustomerId: 11,
		IsActive:   true,
		OrderDate:  &date.Date{Year: 2021, Month: 1, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "CocaCola", ProductType: product.ProductType_DRINK},
		},
	}
	bytes, err := json.Marshal(&orderItem)
	if err != nil {
		log.Fatal("deserialization error:", err)
	}

	db := database.NewDB()

	response := db.AddOrder(&orderItem)
	response2 := db.AddOrder(&orders.Order{
		OrderId:    100,
		CustomerId: 11,
		IsActive:   true,
		OrderDate:  &date.Date{Year: 2021, Month: 1, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "CocaCola", ProductType: product.ProductType_DRINK},
		},
	})

	fmt.Println(response)
	fmt.Println(response2)

	fmt.Println(string(bytes))

	allOrders := db.GetAllOrders()

	for _, order := range allOrders {
		fmt.Println(order.OrderId)
	}

}
