package logic

import (
	"context"

	"go-zero-mall/common/errx"
	"go-zero-mall/service/user/model"
	"go-zero-mall/service/user/rpc/internal/svc"
	"go-zero-mall/service/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserinfoLogic {
	return &UpdateUserinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserinfoLogic) UpdateUserinfo(in *user.UpdateInfoRequest) (*user.UserInfoResponse, error) {
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil && err != model.ErrNotFound{
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "UpdateUserinfo FindOne db err : %v, in : %+v",err, in)
	}
	if res == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.USER_NOTEXIST_ERROR), "errx.USER_NOTEXIST_ERROR err : %v, in : %+v",err, in)
	}

	if in.Name != "" {
		res.Name = in.Name
	}

	if in.Gender == 0 || in.Gender == 1 {
		res.Gender = uint64(in.Gender)
	}
	if in.Mobile != "" {
		res.Mobile = in.Mobile
	}
	if in.Password != "" {
		res.Password = in.Password
	}

	if err := l.svcCtx.UserModel.Update(l.ctx, res); err != nil {
		return nil, errors.Wrapf(errx.NewErrMsg("更新用户信息失败"), "UpdateUserinfo Update err : %v, in : %+v",err, in)
	}
	

	return &user.UserInfoResponse{
		Id: int64(res.Id),
		Name: res.Name,
		Gender: int64(res.Gender),
		Mobile: res.Mobile,
		Password: res.Password,
	}, nil
}
