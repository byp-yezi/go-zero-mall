package logic

import (
	"context"

	"go-zero-mall/service/pay/api/internal/svc"
	"go-zero-mall/service/pay/api/internal/types"
	"go-zero-mall/service/pay/rpc/pb/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 支付创建
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	res, err := l.svcCtx.PayRpc.AddPay(l.ctx, &pay.AddPayReq{
		Uid: req.Uid,
		Oid: req.Oid,
		Amount: req.Amount,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
