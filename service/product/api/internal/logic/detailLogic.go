package logic

import (
	"context"

	"go-zero-mall/service/product/api/internal/svc"
	"go-zero-mall/service/product/api/internal/types"
	"go-zero-mall/service/product/rpc/pb/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 产品详情
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	res, err := l.svcCtx.ProductRpc.GetProductById(l.ctx, &product.GetProductByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id: res.Product.Id,
		Name: res.Product.Name,
		Desc: res.Product.Desc,
		Stock: res.Product.Stock,
		Amount: res.Product.Amount,
		Status: res.Product.Status,
	}, nil
}
