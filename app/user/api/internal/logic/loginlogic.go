package logic

import (
	"context"

	"PanPan/app/user/api/internal/svc"
	"PanPan/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"	
	"PanPan/app/user/rpc/types/user"
	"PanPan/common/errorx"
	"PanPan/utils"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"fmt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户密码形式登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.TokenResp, err error) {
	err = utils.DefaultGetValidParams(l.ctx, req)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(100001, fmt.Sprintf("validate校验错误: %v", err)), "validate校验错误err :%v", err)
	}
	
	cnt, err := l.svcCtx.Rpc.Login(l.ctx, &user.LoginReq{
		PhoneOrEmail: req.PhoneOrEmail,
		PassWord:     req.PassWord,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	accessTokenString, refreshTokenString := utils.GetToken(cnt.UserId, uuid.New().String())
	if accessTokenString == "" || refreshTokenString == "" {
		return nil, errors.Wrapf(errorx.NewCodeError(100002, errorx.JWt), "生成jwt错误 err:%v", err)
	}

	return &types.TokenResp{
		UserId:       cnt.UserId,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil

}
