package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/order/rpc/pb/order"
	"go-zero-mall/service/pay/model"
	"go-zero-mall/service/pay/rpc/internal/svc"
	"go-zero-mall/service/pay/rpc/pb/pay"
	"go-zero-mall/service/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddPayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPayLogic {
	return &AddPayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------pay-----------------------
func (l *AddPayLogic) AddPay(in *pay.AddPayReq) (*pay.AddPayResp, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.USER_NOTEXIST_ERROR), "AddPay UserInfo 用户不存在 err:%v", err)
	}

	// 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.GetOrderById(l.ctx, &order.GetOrderByIdReq{
		Id: in.Oid,
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ORDER_NOTEXIST_ERROR), "AddPay find db 订单不存在 err:%v", err)
	}

	// 查询订单是否创建支付
	resPay, err := l.svcCtx.PayModel.FindOneByOid(l.ctx, in.Oid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "AddPay FindOneByOid err:%v", err)
	}
	if resPay != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ORDER_ALREADY_CREATE_FOR_PAY_ERROR), "AddPay FindOneByOid 已创建过付款订单 err:%v", err)
	}

	newPay := model.Pay {
		Uid: uint64(in.Uid),
		Oid: uint64(in.Oid),
		Amount: uint64(in.Amount),
		Source: 0,
		Status: 0,
	}

	res, err := l.svcCtx.PayModel.Insert(l.ctx, &newPay)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.CREATE_PAY_ERROR), "AddPay Insert 支付创建失败 err:%v", err)
	}

	newPayId, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrMsg("支付id赋值失败"), "AddPay LastInsertId 支付id赋值失败 err:%v", err)
	}
	// newPay.Id = uint64(newPayId)

	return &pay.AddPayResp{
		Id: newPayId,
	}, nil
}
