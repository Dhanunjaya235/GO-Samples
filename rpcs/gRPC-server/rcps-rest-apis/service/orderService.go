package service

import (
	"context"
	"rpc/grpc/examples/protogen/golang/orders"
)

type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

func NewOrderService(db *DB) OrderService {
	return OrderService{
		db: db,
	}
}

func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {

	err := o.db.AddOrder(req.GetOrder())

	return &orders.Empty{}, err

}

func (o *OrderService) GetAllOrders(_ context.Context, req *orders.PayloadWithSingleOrder) ([]*orders.Order, error) {

	orders, err := o.db.GetAllOrders()

	return orders, err

}
