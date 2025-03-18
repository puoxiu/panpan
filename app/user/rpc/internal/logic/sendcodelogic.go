package logic

import (
	"context"

	"PanPan/app/user/rpc/internal/svc"
	"PanPan/app/user/rpc/types/user"
	"PanPan/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *user.SendCodeReq) (*user.SendCodeResp, error) {
	// todo: add your logic here and delete this line

	vecode := utils.SMSV1(in.UserPhone, l.svcCtx.Config.Credential.SecretId, l.svcCtx.Config.Credential.SecretKey, l.ctx, l.svcCtx.Rdb)
	// vecode := utils.SMS(in.UserPhone, l.svcCtx.Config.Credential.SecretId, l.svcCtx.Config.Credential.SecretKey, l.ctx, l.svcCtx.Rdb)

	return &user.SendCodeResp{
		VeCode: vecode,
	}, nil
}
