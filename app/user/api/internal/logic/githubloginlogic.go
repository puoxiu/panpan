package logic

import (
	"context"

	"PanPan/app/user/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GithubLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// github第三方登陆
func NewGithubLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GithubLoginLogic {
	return &GithubLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GithubLoginLogic) GithubLogin() error {
	// todo: add your logic here and delete this line

	return nil
}
