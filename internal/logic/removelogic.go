package logic

import (
	"context"
	"fmt"

	"gitlab.energy-envision.com/storage/client"
	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveLogic {
	return RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) RemoveObject(req types.RemoveReq) error {
	l.Logger.Infof("Remove.receive req:%+v", req)

	claims, err := types.ParseToken(req.S3Authorization)
	if err != nil {
		l.Logger.Errorf("身份验证失败：%s", err.Error())
		return fmt.Errorf("身份验证失败,%s", err.Error())
	}
	err = claims.Valid()
	if err != nil {
		l.Logger.Errorf("身份验证失败：%s", err.Error())
		return fmt.Errorf("身份验证失败,%s", err.Error())
	}

	token, err := types.GetTokenDataFromJwtClaims(claims)
	if err != nil {
		l.Logger.Errorf("身份读取失败：%s", err.Error())
		return fmt.Errorf("身份读取失败,%s", err.Error())
	}

	user, err := l.svcCtx.UserModel.FindOneByAccessKey(token["accessKey"].(string))
	if err != nil {
		l.Logger.Errorf("获取用户信息失败：%s", err.Error())
		return fmt.Errorf("获取用户信息失败,%s", err.Error())
	}

	err = l.svcCtx.Client.RemoveObject(context.Background(), types.BucketName(user.TenantId, user.Namespace),
		types.BucketUserObjectPath(token["accessKey"].(string), token["path"].(string), req.ObjectID), client.RemoveObjectOptions{})
	if err != nil {
		l.Logger.Errorf("删除文件失败：%s", err.Error())
		return fmt.Errorf("删除文件失败,%s", err.Error())
	}

	return nil
}
