package logic

import (
	"context"
	"database/sql"

	"go-zero-mall/service/order/model"
	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb/order"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// 获取RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// 开启子事务屏障
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// // 查询用户是否存在
		// _, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		// 	Id: in.Uid,
		// })
		// if err != nil {
		// 	return fmt.Errorf("用户不存在111")
		// 	// return status.Error(500, err.Error())
		// }

		newOrder := model.Order{
			Uid:    uint64(in.Uid),
			Pid:    uint64(in.Pid),
			Amount: uint64(in.Amount),
			Status: 0,
		}

		// 创建订单
		_, err = l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		return nil
	}); err != nil {
		// return nil, errors.Wrapf(errx.NewErrCode(errx.USER_NOTEXIST_ERROR), "AddOrderRevert CallWithDB UserInfo 用户不存在2, err:%v", err)
		return nil, status.Error(500, err.Error())
	}

	// // 查询产品是否存在
	// resProduct, err := l.svcCtx.ProductRpc.GetProductById(l.ctx, &product.GetProductByIdReq{
	// 	Id: in.Pid,
	// })
	// if err != nil {
	// 	return nil, errors.Wrapf(errx.NewErrCode(errx.GET_PRODUCT_BY_ID_ERROR), "AddOrder GetProductById 产品不存在 pid:%v err:%v", in.Pid, err)
	// }

	// // 判断产品库存是否充足
	// if resProduct.Product.Stock <= 0 {
	// 	return nil, errors.Wrapf(errx.NewErrCode(errx.STOCK_INSUFFICIENT_ERROR), "AddOrder GetProductById 产品库存不足 err:%v", err)
	// }

	// // 更新产品库存
	// _, err = l.svcCtx.ProductRpc.UpdateProduct(l.ctx, &product.UpdateProductReq{
	// 	Id: resProduct.Product.Id,
	// 	Name: resProduct.Product.Name,
	// 	Desc: resProduct.Product.Desc,
	// 	Stock: resProduct.Product.Stock - 1,
	// 	Amount: resProduct.Product.Amount,
	// 	Status: resProduct.Product.Status,
	// })
	// if err != nil {
	// 	return nil, errors.Wrapf(errx.NewErrCode(errx.STOCK_REDUCE_ERROR), "AddOrder UpdateProduct 产品库存减少失败 err:%v", err)
	// }

	return &order.AddOrderResp{}, nil
}
