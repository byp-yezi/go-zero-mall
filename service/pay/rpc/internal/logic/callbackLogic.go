package logic

import (
	"context"

	"go-zero-mall/service/pay/rpc/internal/svc"
	"go-zero-mall/service/pay/rpc/pb/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackReq) (*pay.CallbackResp, error) {
	// todo: add your logic here and delete this line

	return &pay.CallbackResp{}, nil
}
