package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb/order"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderLogic) UpdateOrder(in *order.UpdateOrderReq) (*order.UpdateOrderResp, error) {
	// 查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "UpdateOrder find db err, id:%d, err:%v", in.Id, err)
	}
	if res == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ORDER_NOTEXIST_ERROR), "UpdateOrder FindOne 订单不存在 id:%d, err:%v", in.Id, err)
	}

	var dbOrder = new(model.Order)
	err = copier.Copy(dbOrder, in)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.SERVER_COMMON_ERROR), "UpdateProduct copier id:%d, err:%v", in.Id, err)
	}
	err = l.svcCtx.OrderModel.Update(l.ctx, dbOrder)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ORDER_UPDATE_ERROR), "UpdateOrder Update 订单更新失败 id:%d, err:%v", in.Id, err)
	}

	return &order.UpdateOrderResp{}, nil
}
