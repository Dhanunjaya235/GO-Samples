package main

import (
	"context"
	"fmt"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"rpc/grpc/examples/protogen/golang/orders"
	product "rpc/grpc/examples/protogen/golang/products"
)

func main() {
	conn, err := grpc.Dial("localhost:2035", grpc.WithTransportCredentials(insecure.NewCredentials()))

	orderItem := orders.Order{
		OrderId:    100,
		CustomerId: 11,
		IsActive:   true,
		OrderDate:  &date.Date{Year: 2021, Month: 1, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "CocaCola", ProductType: product.ProductType_DRINK},
		},
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	defer conn.Close()

	client := orders.NewOrdersClient(conn)

	payload := orders.PayloadWithSingleOrder{Order: &orderItem}

	_, error := client.AddOrder(context.Background(), &payload)

	if error != nil {
		fmt.Println(error.Error())
	}

}
