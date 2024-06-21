package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/product/rpc/internal/svc"
	"go-zero-mall/service/product/rpc/pb/product"
	"go-zero-mall/service/user/model"

	"github.com/pkg/errors"
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
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "DelProduct find db err, id:%d, err:%v", in.Id, err)
	}
	if res == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.GET_PRODUCT_BY_ID_ERROR), "id:%d", in.Id)
	}
	err = l.svcCtx.ProductModel.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "DelProduct delete db err, id:%d, err:%v", res.Id, err)
	}

	return &product.DelProductResp{}, nil
}
