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

type PaidOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaidOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaidOrderLogic {
	return &PaidOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaidOrderLogic) PaidOrder(in *order.PaidReq) (*order.PaidResp, error) {
	// 查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "PaidOrder find db err, id:%d, err:%v", in.Id, err)
	}
	if res == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ORDER_NOTEXIST_ERROR), "PaidOrder FindOne 订单不存在 id:%v, err:%v", in.Id, err)
	}

	res.Status = 1

	err = l.svcCtx.OrderModel.Update(l.ctx, res)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ORDER_UPDATE_ERROR), "PaidOrder Update 订单更新失败 id:%d, err:%v", in.Id, err)
	}

	return &order.PaidResp{}, nil
}
