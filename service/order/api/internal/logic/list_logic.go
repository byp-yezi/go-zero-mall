package logic

import (
	"context"

	"go-zero-mall/service/order/api/internal/svc"
	"go-zero-mall/service/order/api/internal/types"
	"go-zero-mall/service/order/rpc/pb/order"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单列表
func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	res, err := l.svcCtx.OrderRpc.ListOrder(l.ctx, &order.ListRequest{
		Uid: req.Uid,
	})
	if err != nil {
		return nil, err
	}

	var orderList []types.Order
	for _, v := range res.Order {
		var resOrder types.Order
		_ = copier.Copy(&resOrder, v)
		orderList = append(orderList, resOrder)
	}

	return &types.ListResponse{
		Orders: orderList,
	}, nil
}
