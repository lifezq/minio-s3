package logic

import (
	"context"

	"github.com/lifezq/minio-s3/internal/svc"
	"github.com/lifezq/minio-s3/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ObjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) ObjectLogic {
	return ObjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ObjectLogic) Object() (*types.UploadResp, error) {
	// todo: add your logic here and delete this line

	return &types.UploadResp{}, nil
}
