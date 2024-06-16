package logic

import (
	"context"

	"go-zero-mall/common/cryptx"
	"go-zero-mall/common/errx"
	"go-zero-mall/service/user/model"
	"go-zero-mall/service/user/rpc/internal/code"
	"go-zero-mall/service/user/rpc/internal/svc"
	"go-zero-mall/service/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	if len(in.Name) == 0 {
		return nil, code.RegisterNameEmpty
	}
	// 判断手机号是否已经注册
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err == nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.USER_EXIST_ERROR), "用户已经存在 mobile:%s, err:%v", in.Mobile, err)
	}
	if err == model.ErrNotFound {
		newUser := model.User{
			Name:     in.Name,
			Gender:   uint64(in.Gender),
			Mobile:   in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		newUserId, err := res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		newUser.Id = uint64(newUserId)
		return &user.RegisterResponse{
			Id:     int64(newUser.Id),
			Name:   newUser.Name,
			Gender: int64(newUser.Gender),
			Mobile: newUser.Mobile,
		}, nil
	}

	return nil, status.Error(500, err.Error())
}
