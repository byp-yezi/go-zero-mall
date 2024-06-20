package user

import (
	"context"
	"encoding/json"

	"go-zero-mall/service/user/api/internal/code"
	"go-zero-mall/service/user/api/internal/svc"
	"go-zero-mall/service/user/api/internal/types"
	"go-zero-mall/service/user/rpc/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户信息
func NewUpdateUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserinfoLogic {
	return &UpdateUserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserinfoLogic) UpdateUserinfo(req *types.UpdateInfoRequest) (resp *types.UserInfoResponse, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, err
	}
	if req.Id != uid {
		return nil, errors.Wrapf(code.ErrUserIdNotExist, "req: %+v, err : %+v", req, err)
	}
	if req.Gender != 0 && req.Gender != 1 {
		return nil, errors.Wrapf(code.ErrGenderOnlyOneOrTwo, "req: %+v", req)
	}

	res, err := l.svcCtx.UserRpc.UpdateUserinfo(l.ctx, &user.UpdateInfoRequest{
		Id: req.Id,
		Name: req.Name,
		Mobile: req.Mobile,
		Gender: req.Gender,
		Password: req.Password,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.UserInfoResponse{
		Id:       res.Id,
		Name:     res.Name,
		Gender:   res.Gender,
		Mobile:   res.Mobile,
		Password: res.Password,
	}, nil
}
