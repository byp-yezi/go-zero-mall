package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/ordersrv"
	"go-zero-mall/service/order/rpc/pb/order"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIdLogic {
	return &GetOrderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByIdLogic) GetOrderById(in *order.GetOrderByIdReq) (*order.GetOrderByIdResp, error) {
	// 查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "GetOrderById find db err, id:%d, err:%v", in.Id, err)
	}
	if res == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ORDER_NOTEXIST_ERROR), "GetOrderById FindOne 订单不存在 id:%v, err:%v", in.Id, err)
	}

	var srvOrder ordersrv.Order
	err = copier.Copy(&srvOrder, res)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.SERVER_COMMON_ERROR), "GetOrderById copier err:%v", err)
	}
	srvOrder.CreateTime = res.CreateTime.Unix()
	srvOrder.UpdateTime = res.UpdateTime.Unix()

	return &order.GetOrderByIdResp{
		Order: &srvOrder,
	}, nil
}
