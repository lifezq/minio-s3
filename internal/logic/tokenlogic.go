package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"
	"golang.org/x/crypto/bcrypt"
)

type TokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokenLogic {
	return TokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenLogic) Token(req types.TokenReq) (*types.TokenResp, error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("Token.receive req:%+v", req)

	user, err := l.svcCtx.UserModel.FindOneByAccessKey(req.AccessKey)
	if err != nil {
		l.Logger.Errorf("获取用户失败：%s\n", err.Error())
		return &types.TokenResp{}, fmt.Errorf("获取用户失败,%s", err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.SecretKey), []byte(req.SecretKey))
	if err != nil {
		l.Logger.Errorf("用户secretKey验证不通过：%s\n", err.Error())
		return &types.TokenResp{}, fmt.Errorf("用户secretKey验证不通过")
	}

	s3Authorization, err := types.CreateToken(req.AccessKey, req.SecretKey, req.Path, l.svcCtx.Config.TokenExpireIn)
	if err != nil {
		l.Logger.Errorf("创建令牌发生错误：%s\n", err.Error())
		return &types.TokenResp{}, fmt.Errorf("创建令牌发生错误")
	}

	return &types.TokenResp{
		S3Authorization: s3Authorization,
		ExpireIn:        l.svcCtx.Config.TokenExpireIn,
	}, nil
}
