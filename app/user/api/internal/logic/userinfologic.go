package logic

import (
	"context"

	"PanPan/app/user/api/internal/svc"
	"PanPan/app/user/api/internal/types"

	"PanPan/app/user/rpc/types/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	cnt, err := l.svcCtx.Rpc.UserInfo(l.ctx, &user.UserInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	ret := types.UserInfoItem{
		Id:         cnt.Users[0].UserId,
		PassWord:   cnt.Users[0].PassWord,
		UserNick:   cnt.Users[0].User_Nick,
		UserFace:   cnt.Users[0].User_Face,
		UserSex:    int64(cnt.Users[0].User_Sex),
		UserEmail:  cnt.Users[0].User_Email,
		UserPhone:  cnt.Users[0].User_Phone,
		CreateTime: cnt.Users[0].CreateTime.AsTime().Format("2006-01-02 15:04:05"),
		UpdateTime: cnt.Users[0].UpdateTime.AsTime().Format("2006-01-02 15:04:05"),
		DeleteTime: cnt.Users[0].DeleteTime.AsTime().Format("2006-01-02 15:04:05"),
	}
	logc.Info(l.ctx, "这里是usrinfo:", ret)
	return &types.UserInfoResp{UserInfo: &ret}, nil
}
