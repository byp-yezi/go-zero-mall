package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/order/api/internal/svc"
	"go-zero-mall/service/order/api/internal/types"
	"go-zero-mall/service/order/rpc/pb/order"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单详情
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	res, err := l.svcCtx.OrderRpc.GetOrderById(l.ctx, &order.GetOrderByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	var orderDetail types.DetailResponse
	if res != nil {
		err = copier.Copy(&orderDetail, res)
		if err != nil {
			return nil, errors.Wrapf(errx.NewErrCode(errx.SERVER_COMMON_ERROR), "GetOrderById copier id:%v err:%v", req.Id, err)
		}
	}

	return &types.DetailResponse{
		Order: orderDetail.Order,
	}, nil
}
