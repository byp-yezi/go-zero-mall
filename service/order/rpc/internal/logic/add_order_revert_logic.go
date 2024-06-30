package logic

import (
	"context"
	"database/sql"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb/order"
	"go-zero-mall/service/user/rpc/user"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type AddOrderRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderRevertLogic {
	return &AddOrderRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddOrderRevertLogic) AddOrderRevert(in *order.AddOrderReq) (*order.AddOrderResp, error) {
	// 获取RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.SERVER_COMMON_ERROR), "AddOrderRevert RawDB 获取RawDB失败, err:%v", err)
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.SERVER_COMMON_ERROR), "AddOrderRevert BarrierFromGrpc 获取Barrier失败, err:%v", err)
	}

	// 开启子事务屏障对象
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 查询用户是否存在
		_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			Id: in.Uid,
		})
		if err != nil {
			return errors.Wrapf(errx.NewErrCode(errx.USER_NOTEXIST_ERROR), "AddOrderRevert CallWithDB UserInfo 用户不存在3, err:%v", err)
		}
		// 查询用户最新创建的订单
		resOrder, err := l.svcCtx.OrderModel.FindOneByUid(l.ctx, in.Uid)
		if err != nil {
			return errors.Wrapf(errx.NewErrCode(errx.ORDER_NOTEXIST_ERROR), "AddOrderRevert CallWithDB FindOneByUid 订单不存在, err:%v", err)
		}
		// 修改订单状态9，标识订单已失效，并更新订单
		resOrder.Status = 9
		err = l.svcCtx.OrderModel.TxUpdate(l.ctx, tx, resOrder)
		if err != nil {
			return errors.Wrapf(errx.NewErrCode(errx.ORDER_UPDATE_ERROR), "AddOrderRevert CallWithDB TxUpdate 订单更新失败, err:%v", err)
		}
		return nil
	}); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.SERVER_COMMON_ERROR), "AddOrderRevert CallWithDB 开启子事务屏障失败, err:%v", err)
	}

	return &order.AddOrderResp{}, nil
}
