package logic

import (
	"context"
	"fmt"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/order/model"
	"go-zero-mall/service/order/rpc/internal/svc"
	"go-zero-mall/service/order/rpc/pb/order"
	"go-zero-mall/service/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type ListOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrderLogic {
	return &ListOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListOrderLogic) ListOrder(in *order.ListRequest) (*order.ListResp, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.USER_NOTEXIST_ERROR), "ListOrder UserInfo 用户不存在 uid:%v err:%v", in.Uid, err)
	}

	// 查询订单是否存在
	list, err := l.svcCtx.OrderModel.FindAllByUid(l.ctx, in.Uid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "FindAllByUid find db err, uid:%d, err:%v", in.Uid, err)
	}
	if list == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.ORDER_NOTEXIST_ERROR), "FindAllByUid FindOne 订单不存在 uid:%v, err:%v", in.Uid, err)
	}

	var resp []*order.Order
	if len(list) > 0 {
		mr.MapReduceVoid(func(source chan<- interface{}) {
			for _, order := range list {
				fmt.Println("order.Id=====================", order.Id)
				source <- order.Id
			}
		}, func(item interface{}, writer mr.Writer[*model.Order], cancel func(error)) {
			id := item.(uint64)
			orderResp, err := l.svcCtx.OrderModel.FindOne(l.ctx, uint64(id))
			if err != nil && err != model.ErrNotFound {
				cancel(err)
			}
			if orderResp != nil && orderResp.Id > 0 {
				writer.Write(orderResp)
			}
		}, func(pipe <-chan *model.Order, cancel func(error)) {
			for orderResp2 := range pipe {
				var tempOrder order.Order
				_ = copier.Copy(&tempOrder, &orderResp2)
				tempOrder.CreateTime = orderResp2.CreateTime.Unix()
				tempOrder.UpdateTime = orderResp2.UpdateTime.Unix()
				resp = append(resp, &tempOrder)
			}
		})
	}

	// orderList := make([]*order.Order, 0)
	// for _, item := range list {
	// 	orderList = append(orderList, &order.Order{
	// 		Id: int64(item.Id),
	// 		Uid: int64(item.Uid),
	// 		Pid: int64(item.Pid),
	// 		Amount: int64(item.Amount),
	// 		Status: int64(item.Status),
	// 		CreateTime: item.CreateTime.Unix(),
	// 		UpdateTime: item.UpdateTime.Unix(),
	// 	})
	// }

	// return &order.ListResp{
	// 	Order: orderList,
	// }, nil
	return &order.ListResp{
		Order: resp,
	}, nil
}
