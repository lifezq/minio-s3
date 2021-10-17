package logic

import (
	"context"
	"log"

	"minio-s3/internal/svc"
	"minio-s3/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) UploadLogic {
	return UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req types.UploadReq) (*types.UploadResp, error) {
	// todo: add your logic here and delete this line

	log.Printf("upload filename:%d\n", len(req.Filename))
	return &types.UploadResp{}, nil
}
