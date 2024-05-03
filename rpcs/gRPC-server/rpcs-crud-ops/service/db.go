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
	for _, o := range d.collection {
		if o.OrderId == order.OrderId {
			fmt.Println("Duplicate Order")
			return fmt.Errorf("duplicate Order : %d", order.OrderId)
		}
	}

	d.collection = append(d.collection, order)

	return nil

}

func (d *DB) GetAllOrders(payload *orders.Empty) ([]*orders.Order, error) {

	return d.collection, nil
}

func (d *DB) GetOrder(payload *orders.PayloadWithOrderID) (*orders.Order, error) {

	for _, o := range d.collection {
		if o.OrderId == payload.GetOrderId() {
			return o, nil
		}
	}

	err := fmt.Errorf("not found order with id %d", payload.OrderId)
	return nil, err
}
