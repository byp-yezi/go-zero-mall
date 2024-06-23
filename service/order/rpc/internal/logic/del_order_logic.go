package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb/order"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DelOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOrderLogic {
	return &DelOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOrderLogic) DelOrder(in *order.DelOrderReq) (*order.DelOrderResp, error) {
	// 查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "DelOrder FindOne id:%v, err:%v", in.Id, err)
	}
	if res == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ORDER_NOTEXIST_ERROR), "DelOrder FindOne 订单不存在 id:%v, err:%v", in.Id, err)
	}
	err = l.svcCtx.OrderModel.Delete(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ORDER_DELETE_ERROR), "DelOrder Delete 订单删除失败 id:%v, err:%v", in.Id, err)
	}

	return &order.DelOrderResp{}, nil
}
