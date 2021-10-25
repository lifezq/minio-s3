package logic

import (
	"context"
	"fmt"
	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"
	"gitlab.energy-envision.com/storage/utils"
	"net/url"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) DownloadLogic {
	return DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req types.DownloadReq) (*types.DownloadResp, error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("Download.receive req:%+v", req)

	claims, err := types.ParseToken(req.S3Authorization)
	if err != nil {
		l.Logger.Errorf("身份验证失败：%s\n", err.Error())
		return nil, fmt.Errorf("身份验证失败,%s", err.Error())
	}
	err = claims.Valid()
	if err != nil {
		l.Logger.Errorf("身份验证失败：%s\n", err.Error())
		return nil, fmt.Errorf("身份验证失败,%s", err.Error())
	}

	token, err := types.GetTokenDataFromJwtClaims(claims)
	if err != nil {
		l.Logger.Errorf("身份读取失败：%s\n", err.Error())
		return nil, fmt.Errorf("身份读取失败,%s", err.Error())
	}

	user, err := l.svcCtx.UserModel.FindOneByAccessKey(token["accessKey"].(string))
	if err != nil {
		l.Logger.Errorf("获取用户失败：%s\n", err.Error())
		return nil, fmt.Errorf("获取用户失败,%s", err.Error())
	}

	ctx := context.Background()
	downUrl, err := l.svcCtx.Client.PresignedGetObject(ctx, req.Bucket, fmt.Sprintf("%s/%s/%s",
		types.BucketHome(user.AccessKey), utils.PathFilter(req.Path),
		req.ObjectID), time.Hour, url.Values{})
	if err != nil {
		l.Logger.Errorf("访问文件失败：%s\n", err.Error())
		return nil, fmt.Errorf("访问文件失败,%s", err.Error())
	}

	return &types.DownloadResp{
		Url: downUrl.String(),
	}, nil
}
