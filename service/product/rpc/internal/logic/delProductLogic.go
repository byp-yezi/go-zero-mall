package logic

import (
	"context"

	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProductLogic {
	return &DelProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelProductLogic) DelProduct(in *product.DelProductReq) (*product.DelProductResp, error) {
	// todo: add your logic here and delete this line

	return &product.DelProductResp{}, nil
}
