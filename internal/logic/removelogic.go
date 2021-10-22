package logic

import (
	"context"
	"fmt"

	"github.com/lifezq/minio-s3/client"
	"github.com/lifezq/minio-s3/internal/svc"
	"github.com/lifezq/minio-s3/internal/types"

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
	l.Logger.Infof("Remove.receive req:%+v\n", req)

	var token types.S3AuthorizationToken
	err := l.svcCtx.Cache.Get(types.CacheS3AuthorizationKey(req.S3Authorization), &token)
	if err != nil {
		l.Logger.Errorf("身份验证失败：%s\n", err.Error())
		return fmt.Errorf("身份验证失败,%s", err.Error())
	}

	user, err := l.svcCtx.UserModel.FindOneByAccessKey(token.AccessKey)
	if err != nil {
		l.Logger.Errorf("获取用户失败：%s\n", err.Error())
		return fmt.Errorf("获取用户失败,%s", err.Error())
	}

	l.svcCtx.Client.RemoveObject(context.Background(), types.BucketName(user.TenantId, user.Namespace),
		types.BucketUserObjectPath(token.AccessKey, token.Path, req.ObjectID), client.RemoveObjectOptions{})
	return nil
}
