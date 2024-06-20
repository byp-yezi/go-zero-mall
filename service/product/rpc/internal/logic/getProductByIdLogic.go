package logic

import (
	"context"

	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductByIdLogic {
	return &GetProductByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductByIdLogic) GetProductById(in *product.GetProductByIdReq) (*product.GetProductByIdResp, error) {
	// todo: add your logic here and delete this line

	return &product.GetProductByIdResp{}, nil
}
