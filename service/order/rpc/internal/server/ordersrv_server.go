// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"go-zero-mall/service/order/rpc/internal/logic"
	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb/order"
)

type OrdersrvServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrdersrvServer
}

func NewOrdersrvServer(svcCtx *svc.ServiceContext) *OrdersrvServer {
	return &OrdersrvServer{
		svcCtx: svcCtx,
	}
}

// -----------------------order-----------------------
func (s *OrdersrvServer) AddOrder(ctx context.Context, in *order.AddOrderReq) (*order.AddOrderResp, error) {
	l := logic.NewAddOrderLogic(ctx, s.svcCtx)
	return l.AddOrder(in)
}

func (s *OrdersrvServer) UpdateOrder(ctx context.Context, in *order.UpdateOrderReq) (*order.UpdateOrderResp, error) {
	l := logic.NewUpdateOrderLogic(ctx, s.svcCtx)
	return l.UpdateOrder(in)
}

func (s *OrdersrvServer) DelOrder(ctx context.Context, in *order.DelOrderReq) (*order.DelOrderResp, error) {
	l := logic.NewDelOrderLogic(ctx, s.svcCtx)
	return l.DelOrder(in)
}

func (s *OrdersrvServer) GetOrderById(ctx context.Context, in *order.GetOrderByIdReq) (*order.GetOrderByIdResp, error) {
	l := logic.NewGetOrderByIdLogic(ctx, s.svcCtx)
	return l.GetOrderById(in)
}

func (s *OrdersrvServer) SearchOrder(ctx context.Context, in *order.SearchOrderReq) (*order.SearchOrderResp, error) {
	l := logic.NewSearchOrderLogic(ctx, s.svcCtx)
	return l.SearchOrder(in)
}

func (s *OrdersrvServer) ListOrder(ctx context.Context, in *order.ListRequest) (*order.ListResp, error) {
	l := logic.NewListOrderLogic(ctx, s.svcCtx)
	return l.ListOrder(in)
}

func (s *OrdersrvServer) PaidOrder(ctx context.Context, in *order.PaidReq) (*order.PaidResp, error) {
	l := logic.NewPaidOrderLogic(ctx, s.svcCtx)
	return l.PaidOrder(in)
}
