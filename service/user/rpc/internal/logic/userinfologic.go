package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/user/model"
	"go-zero-mall/service/user/rpc/internal/svc"
	"go-zero-mall/service/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(errx.NewErrCode(errx.USER_NOTEXIST_ERROR), "err:%v", err)
		}
		return nil, status.Error(500, err.Error())
	}

	return &user.UserInfoResponse{
		Id:       int64(res.Id),
		Name:     res.Name,
		Gender:   int64(res.Gender),
		Mobile:   res.Mobile,
		Password: res.Password,
	}, nil
}
