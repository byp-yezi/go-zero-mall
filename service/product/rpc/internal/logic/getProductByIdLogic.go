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
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "GetProductById find db err, id:%d, err:%v", in.Id, err)
	}
	if res == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.GET_PRODUCT_BY_ID_ERROR), "id:%d", in.Id)
	}
	var resProduct productservice.Product
	err = copier.Copy(&resProduct, res)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.SERVER_COMMON_ERROR), "GetProductById copier err:%v", err)
	}
	resProduct.CreateTime = res.CreateTime.Unix()
	resProduct.UpdateTime = res.UpdateTime.Unix()

	return &product.GetProductByIdResp{
		Product: &resProduct,
	}, nil
}
