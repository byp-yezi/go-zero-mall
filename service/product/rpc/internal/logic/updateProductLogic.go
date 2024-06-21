package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/product/model"
	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/pb/product"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *product.UpdateProductReq) (*product.UpdateProductResp, error) {
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "UpdateProduct find db err, id:%d, err:%v", in.Id, err)
	}
	if res == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.GET_PRODUCT_BY_ID_ERROR), "id:%d", in.Id)
	}
	dbProduct := new(model.Product)
	err = copier.Copy(dbProduct, in)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.SERVER_COMMON_ERROR), "UpdateProduct copier err:%v", err)
	}
	err = l.svcCtx.ProductModel.Update(l.ctx, dbProduct)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "UpdateProduct find db err, id:%d, err:%v", in.Id, err)
	}

	return &product.UpdateProductResp{}, nil
}
