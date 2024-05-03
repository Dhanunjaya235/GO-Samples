package database

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

	for _, o := range d.collection {
		if o.OrderId == order.OrderId {
			return fmt.Errorf("duplicate Order : %d", order.OrderId)
		}
	}

	d.collection = append(d.collection, order)

	return nil

}

func (d *DB) GetAllOrders() []*orders.Order {

	return d.collection
}
