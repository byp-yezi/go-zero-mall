package logic

import (
	"context"

	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOrderLogic {
	return &SearchOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOrderLogic) SearchOrder(in *order.SearchOrderReq) (*order.SearchOrderResp, error) {
	// todo: add your logic here and delete this line

	return &order.SearchOrderResp{}, nil
}
