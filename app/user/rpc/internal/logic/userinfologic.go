package logic

import (
	"context"

	"PanPan/app/user/model"
	"PanPan/app/user/rpc/internal/svc"
	"PanPan/app/user/rpc/types/user"
	"PanPan/common/errorx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserList, error) {

	user0 := model.User{}
	r := l.svcCtx.MasterDb.Where("user_id = ?", in.UserId).First(&user0)
	if r.RowsAffected == 0 {
		return nil, errors.Wrapf(errorx.NewCodeError(20001, errorx.ErrMysqlDateNoResult), "用户信息数据库查询为空")
	}
	if r.Error != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(r.Error.Error()), "用户信息数据库查询错误 err：%v", r.Error)
	}
	users := make([]*user.UserInfo, 0)
	user1 := &user.UserInfo{
		UserId:     user0.UserId,
		PassWord:   user0.PassWord,
		User_Nick:  user0.UserNick,
		User_Face:  user0.UserFace,
		User_Sex:   user0.UserSex,
		User_Email: user0.UserEmail,
		User_Phone: user0.UserPhone,
		CreateTime: timestamppb.New(user0.CreateTime),
		UpdateTime: timestamppb.New(user0.UpdateTime),
		DeleteTime: timestamppb.New(user0.DeleteTime.Time),
	}
	users = append(users, user1)
	logc.Info(l.ctx, "这里是users:   ", users)

	return &user.UserList{
		Users: users,
	}, nil
}