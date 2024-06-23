package logic

import (
	"context"

	"go-zero-mall/service/product/api/internal/svc"
	"go-zero-mall/service/product/api/internal/types"
	"go-zero-mall/service/product/rpc/pb/product"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 产品列表
func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	res, err := l.svcCtx.ProductRpc.SearchProduct(l.ctx, &product.SearchProductReq{
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}
	var products []types.Product
	for _, v := range res.Product {
		var resProduct types.Product
		_ = copier.Copy(&resProduct, v)
		products = append(products, resProduct)
	}

	return &types.ListResponse{
		Products: products,
	}, nil
}
