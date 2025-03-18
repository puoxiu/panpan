package logic

import (
	"context"

	"PanPan/app/user/rpc/internal/svc"
	"PanPan/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/pkg/errors"
	"PanPan/app/user/model"
	"PanPan/common/errorx"
	"PanPan/utils"
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

func (l *LoginLogic) Login(in *user.LoginReq) (*user.CommonResp, error) {
	user0 := model.User{}
	r := l.svcCtx.MasterDb.Where("user_phone = ? or user_email = ?", in.PhoneOrEmail, in.PhoneOrEmail).First(&user0)

	if r.RowsAffected == 0 {
		return nil, errors.Wrapf(errorx.NewCodeError(10005, errorx.ERRPhoneOrEmail), "mobile:%s,phone:%v", in.PhoneOrEmail, in.PhoneOrEmail)
	}
	if r.Error != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(r.Error.Error()), "mobile:%s,phone:%v", in.PhoneOrEmail, in.PhoneOrEmail)

	}
	if !utils.CheckPasswordHash(in.PassWord, user0.PassWord) {
		return nil, errors.Wrapf(errorx.NewCodeError(10006, errorx.ERRLoginPassword), "password:%v", in.PassWord)
	}
	return &user.CommonResp{
		UserId: user0.UserId,
	}, nil
}