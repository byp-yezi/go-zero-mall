package logic

import (
	"context"

	"go-zero-mall/service/pay/rpc/internal/svc"
	"go-zero-mall/service/pay/rpc/pb/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPayByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPayByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPayByIdLogic {
	return &GetPayByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPayByIdLogic) GetPayById(in *pay.DetailReq) (*pay.DetailResp, error) {
	// todo: add your logic here and delete this line

	return &pay.DetailResp{}, nil
}
