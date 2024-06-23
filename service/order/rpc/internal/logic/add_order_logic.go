package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb/order"
	"go-zero-mall/service/product/rpc/pb/product"
	"go-zero-mall/service/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------order-----------------------
func (l *AddOrderLogic) AddOrder(in *order.AddOrderReq) (*order.AddOrderResp, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.USER_NOTEXIST_ERROR), "AddOrder UserInfo 用户不存在 uid:%v err:%v", in.Uid, err)
	}

	// 查询产品是否存在
	resProduct, err := l.svcCtx.ProductRpc.GetProductById(l.ctx, &product.GetProductByIdReq{
		Id: in.Pid,
	})
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.GET_PRODUCT_BY_ID_ERROR), "AddOrder GetProductById 产品不存在 pid:%v err:%v", in.Pid, err)
	}

	// 判断产品库存是否充足
	if resProduct.Product.Stock <= 0 {
		return nil, errors.Wrapf(errx.NewErrCode(errx.STOCK_INSUFFICIENT_ERROR), "AddOrder GetProductById 产品库存不足 err:%v", err)
	}

	newOrder := model.Order{
		Uid: uint64(in.Uid),
		Pid: uint64(in.Pid),
		Amount: uint64(in.Amount),
		Status: 0,
	}

	// 创建订单
	res, err := l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.CREATE_ORDER_ERROR), "AddOrder Insert 订单创建失败 err:%v", err)
	}
	newOrderId, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrMsg("订单id赋值失败"), "AddOrder LastInsertId 订单id赋值失败 err:%v", err)
	}
	newOrder.Id = uint64(newOrderId)
	// 更新产品库存
	_, err = l.svcCtx.ProductRpc.UpdateProduct(l.ctx, &product.UpdateProductReq{
		Id: resProduct.Product.Id,
		Name: resProduct.Product.Name,
		Desc: resProduct.Product.Desc,
		Stock: resProduct.Product.Stock - 1,
		Amount: resProduct.Product.Amount,
		Status: resProduct.Product.Status,
	})
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.STOCK_REDUCE_ERROR), "AddOrder UpdateProduct 产品库存减少失败 err:%v", err)
	}

	return &order.AddOrderResp{
		Id: newOrderId,
	}, nil
}
