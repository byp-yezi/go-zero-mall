package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/product/model"
	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/pb/product"
	"go-zero-mall/service/product/rpc/productservice"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchProductLogic) SearchProduct(in *product.SearchProductReq) (*product.SearchProductResp, error) {
	builder := l.svcCtx.ProductModel.SelectBuilder()
	productList, err := l.svcCtx.ProductModel.FindPageListByPage(l.ctx, builder, in.Page, in.Limit, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "SearchProduct find db err, err:%v", err)
	}
	if productList == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.GET_PRODUCT_BY_ID_ERROR), "err:%v", err)
	}

	var resProducts []*productservice.Product
	for _, v := range productList {
		var resProduct productservice.Product
		_ = copier.Copy(&resProduct, &v)
		resProduct.CreateTime = v.CreateTime.Unix()
		resProduct.UpdateTime = v.UpdateTime.Unix()
		resProducts = append(resProducts, &resProduct)
	}

	return &product.SearchProductResp{
		Product: resProducts,
	}, nil
}
