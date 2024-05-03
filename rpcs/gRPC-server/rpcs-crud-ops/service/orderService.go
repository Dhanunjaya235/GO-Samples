package service

import (
	"context"
	"net/http"
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

func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.ResponseFormat, error) {

	err := o.db.AddOrder(req.GetOrder())

	if err != nil {
		return nil, err
	}

	response := orders.ResponseFormat{
		IsSuccess: true,
		Status:    http.StatusOK,
		OrderId:   req.Order.OrderId,
	}

	return &response, err

}

func (o *OrderService) GetAllOrders(_ context.Context, req *orders.Empty) (*orders.PayloadWithMultipelOrders, error) {

	empty := orders.Empty{}

	allOrders, err := o.db.GetAllOrders(&empty)

	response := orders.PayloadWithMultipelOrders{
		Orders: allOrders,
	}

	return &response, err

}

func (o *OrderService) GetOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.PayloadWithSingleOrder, error) {

	order, err := o.db.GetOrder(req)

	response := orders.PayloadWithSingleOrder{
		Order: order,
	}

	return &response, err

}
