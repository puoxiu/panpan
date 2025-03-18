package logic

import (
	"context"

	"PanPan/app/user/rpc/internal/svc"
	"PanPan/app/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"

	"PanPan/app/user/model"
	"PanPan/common/errorx"
	"PanPan/utils"
	"time"

	"github.com/pkg/errors"
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

// Register 用户验证码形式注册/登陆, 在注册之前 需要先获取验证码（此操作会将手机号码和验证码存入redis）
func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.CommonResp, error) {
	vc, err := l.svcCtx.Rdb.Get(l.ctx, in.UserPhone).Result()
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(10003, errorx.ERRNoPhone), "该手机号码不存在: %v", in.UserPhone)
	}
	if in.VeCode != vc {
		return nil, errors.Wrapf(errorx.NewCodeError(10004, errorx.ERRValidateCode), "验证码错误：%v", in.VeCode)
	}
	users, err := l.svcCtx.UserModel.FindUserBy(l.svcCtx.MasterDb, "user_phone", in.UserPhone)
	if err != nil {
		return nil, err
	}
	var user0 model.User
	if len(users) == 0 {
		logc.Info(l.ctx, "该用户为新用户，开始注册")
		// 新建用户
		user0 = model.User{
			PassWord:   utils.HashPassword(utils.GeneratePassword(8)),
			UserNick:   utils.RandNickname(),
			UserSex:    2,
			UserPhone:  in.UserPhone,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		l.svcCtx.MasterDb.Create(&user0)
		return &user.CommonResp{
			UserId: user0.UserId,
		}, nil
	} else {
		user0 = users[0]
		logc.Info(l.ctx, "该用户已经注册，直接登陆")
		return &user.CommonResp{
			UserId: user0.UserId,
		}, nil
	}
}
