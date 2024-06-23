// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package ordersrv

import (
	"context"

	"go-zero-mall/service/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddOrderReq      = order.AddOrderReq
	AddOrderResp     = order.AddOrderResp
	DelOrderReq      = order.DelOrderReq
	DelOrderResp     = order.DelOrderResp
	GetOrderByIdReq  = order.GetOrderByIdReq
	GetOrderByIdResp = order.GetOrderByIdResp
	ListRequest      = order.ListRequest
	ListResp         = order.ListResp
	Order            = order.Order
	PaidReq          = order.PaidReq
	PaidResp         = order.PaidResp
	SearchOrderReq   = order.SearchOrderReq
	SearchOrderResp  = order.SearchOrderResp
	UpdateOrderReq   = order.UpdateOrderReq
	UpdateOrderResp  = order.UpdateOrderResp

	Ordersrv interface {
		// -----------------------order-----------------------
		AddOrder(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderResp, error)
		UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*UpdateOrderResp, error)
		DelOrder(ctx context.Context, in *DelOrderReq, opts ...grpc.CallOption) (*DelOrderResp, error)
		GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*GetOrderByIdResp, error)
		SearchOrder(ctx context.Context, in *SearchOrderReq, opts ...grpc.CallOption) (*SearchOrderResp, error)
		ListOrder(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResp, error)
		PaidOrder(ctx context.Context, in *PaidReq, opts ...grpc.CallOption) (*PaidResp, error)
	}

	defaultOrdersrv struct {
		cli zrpc.Client
	}
)

func NewOrdersrv(cli zrpc.Client) Ordersrv {
	return &defaultOrdersrv{
		cli: cli,
	}
}

// -----------------------order-----------------------
func (m *defaultOrdersrv) AddOrder(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderResp, error) {
	client := order.NewOrdersrvClient(m.cli.Conn())
	return client.AddOrder(ctx, in, opts...)
}

func (m *defaultOrdersrv) UpdateOrder(ctx context.Context, in *UpdateOrderReq, opts ...grpc.CallOption) (*UpdateOrderResp, error) {
	client := order.NewOrdersrvClient(m.cli.Conn())
	return client.UpdateOrder(ctx, in, opts...)
}

func (m *defaultOrdersrv) DelOrder(ctx context.Context, in *DelOrderReq, opts ...grpc.CallOption) (*DelOrderResp, error) {
	client := order.NewOrdersrvClient(m.cli.Conn())
	return client.DelOrder(ctx, in, opts...)
}

func (m *defaultOrdersrv) GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*GetOrderByIdResp, error) {
	client := order.NewOrdersrvClient(m.cli.Conn())
	return client.GetOrderById(ctx, in, opts...)
}

func (m *defaultOrdersrv) SearchOrder(ctx context.Context, in *SearchOrderReq, opts ...grpc.CallOption) (*SearchOrderResp, error) {
	client := order.NewOrdersrvClient(m.cli.Conn())
	return client.SearchOrder(ctx, in, opts...)
}

func (m *defaultOrdersrv) ListOrder(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResp, error) {
	client := order.NewOrdersrvClient(m.cli.Conn())
	return client.ListOrder(ctx, in, opts...)
}

func (m *defaultOrdersrv) PaidOrder(ctx context.Context, in *PaidReq, opts ...grpc.CallOption) (*PaidResp, error) {
	client := order.NewOrdersrvClient(m.cli.Conn())
	return client.PaidOrder(ctx, in, opts...)
}
