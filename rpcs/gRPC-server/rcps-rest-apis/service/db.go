package service

import (
	"fmt"
	"rpc/grpc/examples/protogen/golang/orders"
)

type DB struct {
	collection []*orders.Order
}

func NewDB() *DB {
	return &DB{
		collection: make([]*orders.Order, 0),
	}
}

func (d *DB) AddOrder(order *orders.Order) error {
	fmt.Println("Add Order Called")
	fmt.Println(order.OrderId)
	for _, o := range d.collection {
		if o.OrderId == order.OrderId {
			fmt.Println("Duplicate Order")
			return fmt.Errorf("duplicate Order : %d", order.OrderId)
		}
	}

	d.collection = append(d.collection, order)

	return nil

}

func (d *DB) GetAllOrders() ([]*orders.Order, error) {

	return d.collection, nil
}