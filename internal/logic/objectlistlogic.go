package logic

import (
	"context"

	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ObjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewObjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ObjectListLogic {
	return ObjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ObjectListLogic) ObjectList(req types.ObjectListReq) (*types.ObjectListResp, error) {
	// todo: add your logic here and delete this line

	return &types.ObjectListResp{}, nil
}
