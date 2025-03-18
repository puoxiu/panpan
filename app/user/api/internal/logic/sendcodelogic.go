package logic

import (
	"context"

	"PanPan/app/user/api/internal/svc"
	"PanPan/app/user/api/internal/types"
	"PanPan/app/user/rpc/types/user"
	"PanPan/common/errorx"
	"PanPan/utils"

	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/pkg/errors"
)

type SendcodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发验证码
func NewSendcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendcodeLogic {
	return &SendcodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Sendcode 发送验证码 调用rpc服务
func (l *SendcodeLogic) Sendcode(req *types.RegisterByPhoneRep) (resp *types.RegisterByPhoneResp, err error) {
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errorx.NewCodeError(100001, fmt.Sprintf("validate校验错误: %v", err))
	}

	cnt, err := l.svcCtx.Rpc.SendCode(l.ctx, &user.SendCodeReq{UserPhone: req.UserPhone})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return &types.RegisterByPhoneResp{VeCode: cnt.VeCode}, nil

}
