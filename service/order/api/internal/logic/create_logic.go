package logic

import (
	"context"
	"fmt"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/order/api/internal/svc"
	"go-zero-mall/service/order/api/internal/types"
	"go-zero-mall/service/order/rpc/pb/order"
	"go-zero-mall/service/product/rpc/pb/product"
	"go-zero-mall/service/user/rpc/user"

	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订单创建
func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// res, err := l.svcCtx.OrderRpc.AddOrder(l.ctx, &order.AddOrderReq{
	// 	Uid: req.Uid,
	// 	Pid: req.Pid,
	// 	Amount: req.Amount,
	// 	Status: req.Status,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// 查询用户是否存在
	_, err = l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: req.Uid,
	})
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.USER_NOTEXIST_ERROR), "order-api Create UserInfo 用户不存在222, err:%+v\n", err)
		// return status.Error(500, err.Error())
	}

	// 获取OrderRpc BuildTarget
	orderRpcBusiServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, err
	}

	// 获取ProductRpc BuildTarget
	productRpcBusiServer, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, err
	}

	// dtm 服务的 etcd 注册地址
	var dtmServer = "etcd://192.168.0.111:2379/dtmservice"
	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	// 创建一个saga协议的事务
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcBusiServer+"/order.ordersrv/AddOrder", orderRpcBusiServer+"/order.ordersrv/AddOrderRevert", &order.AddOrderReq{
			Uid:    req.Uid,
			Pid:    req.Pid,
			Amount: req.Amount,
			Status: 0,
		}).
		Add(productRpcBusiServer+"/product.ProductService/DecrStock", productRpcBusiServer+"/product.ProductService/DecrStockRevert", &product.DecrStockReq{
			Id:  req.Pid,
			Num: 1,
		})

	// 事务提交
	err = saga.Submit()
	dtmimp.FatalIfError(err)
	if err != nil {
		return nil, fmt.Errorf("submit data to dtm-server err : %v", err)
	}
	return &types.CreateResponse{}, nil
}
