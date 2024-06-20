package logic

import (
	"context"

	"go-zero-mall/common/cryptx"
	"go-zero-mall/common/errx"
	"go-zero-mall/service/user/model"
	"go-zero-mall/service/user/rpc/internal/svc"
	"go-zero-mall/service/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrapf(errx.NewErrCode(errx.USER_NOTEXIST_ERROR), "Login FindOneByMobile 用户不存在 in : %+v, err : %+v", in, err)
		}
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "Login FindOneByMobile db err : %v, in : %+v", err, in)
	}
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, errors.Wrapf(errx.NewErrCode(errx.PASSWORD_ERROR), "err : %v", err)
	}

	return &user.LoginResponse{
		Id:     int64(res.Id),
		Name:   res.Name,
		Gender: int64(res.Gender),
		Mobile: res.Mobile,
	}, nil
}
