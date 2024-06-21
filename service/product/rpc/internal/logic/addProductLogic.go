package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/product/model"
	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/pb/product"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------product-----------------------
func (l *AddProductLogic) AddProduct(in *product.AddProductReq) (*product.AddProductResp, error) {
	newProduct := model.Product{
		Name: in.Name,
		Desc: in.Desc,
		Stock: uint64(in.Stock),
		Amount: uint64(in.Amount),
		Status: uint64(in.Status),
	}

	res, err := l.svcCtx.ProductModel.Insert(l.ctx, &newProduct)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.CREATE_PRODUCT_ERROR), "err : %+v", err)
	}
	newProductId, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrMsg("产品id赋值失败"), "err : %+v", err)
	}
	newProduct.Id = uint64(newProductId)

	return &product.AddProductResp{}, nil
}
