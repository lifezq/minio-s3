package logic

import (
	"context"
	"fmt"

	"gitlab.energy-envision.com/storage/client"
	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"
	"gitlab.energy-envision.com/storage/utils"

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
	l.Logger.Infof("ObjectList.receive req:%+v", req)

	claims, err := types.ParseToken(req.S3Authorization)
	if err != nil {
		l.Logger.Errorf("身份验证失败：%s", err.Error())
		return nil, fmt.Errorf("身份验证失败,%s", err.Error())
	}
	err = claims.Valid()
	if err != nil {
		l.Logger.Errorf("身份验证失败：%s", err.Error())
		return nil, fmt.Errorf("身份验证失败,%s", err.Error())
	}

	token, err := types.GetTokenDataFromJwtClaims(claims)
	if err != nil {
		l.Logger.Errorf("身份读取失败：%s", err.Error())
		return nil, fmt.Errorf("身份读取失败,%s", err.Error())
	}

	user, err := l.svcCtx.UserModel.FindOneByAccessKey(token["accessKey"].(string))
	if err != nil {
		l.Logger.Errorf("获取用户信息失败：%s", err.Error())
		return nil, fmt.Errorf("获取用户信息失败,%s", err.Error())
	}

	bucketHome := types.BucketHome(user.AccessKey)
	ret := &types.ObjectListResp{}
	for objectInfo := range l.svcCtx.Client.ListObjects(context.Background(),
		types.BucketName(user.TenantId, user.Namespace),
		client.ListObjectsOptions{
			Prefix:       bucketHome + "/" + utils.PathFilter(req.Prefix),
			Recursive:    true,
			MaxKeys:      -1,
			WithMetadata: true,
		}) {

		if len(objectInfo.Key) < len(bucketHome) {
			continue
		}

		meta := make(map[string]string)
		delete(objectInfo.UserMetadata, "content-type")
		for k, v := range objectInfo.UserMetadata {
			meta[types.MetadataStripX(k)] = v
		}

		objectType, path, objectID := types.PathAndObjectIDFromObjectName(objectInfo.Key[len(bucketHome):])
		ret.Items = append(ret.Items, types.ObjectInfo{
			Type:         objectType,
			ObjectID:     objectID,
			Path:         path,
			Size:         objectInfo.Size,
			ETag:         objectInfo.ETag,
			LastModified: objectInfo.LastModified,
			Metadata:     meta,
		})
	}
	return ret, nil
}
