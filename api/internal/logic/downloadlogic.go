package logic

import (
	"context"

	"github.com/lifezq/minio-s3/api/internal/svc"
	"github.com/lifezq/minio-s3/api/internal/types"

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

func (l *DownloadLogic) Download(req types.DownloadReq) error {
	// todo: add your logic here and delete this line

	return nil
}
