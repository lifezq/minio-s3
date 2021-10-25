package logic

import (
	"context"
	"fmt"
	"gitlab.energy-envision.com/storage/utils"
	"net/http"
	"path"
	"time"

	"gitlab.energy-envision.com/storage/client"
	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"

	"github.com/lifezq/goutils/hash"
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

func (l *ObjectLogic) Object(r *http.Request) (*types.UploadResp, error) {
	// todo: add your logic here and delete this line

	claims, err := types.ParseToken(r.Header.Get(types.S3_AUTHORIZATION))
	if err != nil {
		l.Logger.Errorf("身份解析失败，%s", err.Error())
		return nil, fmt.Errorf("身份解析失败")
	}

	err = claims.Valid()
	if err != nil {
		l.Logger.Errorf("身份验证失败，%s", err.Error())
		return nil, fmt.Errorf("身份验证失败")
	}

	token, err := types.GetTokenDataFromJwtClaims(claims)
	if err != nil {
		l.Logger.Errorf("身份读取失败，%s", err.Error())
		return nil, fmt.Errorf("身份读取失败")
	}

	accessKey := token["accessKey"].(string)
	uploadPath := token["path"].(string)

	user, err := l.svcCtx.UserModel.FindOneByAccessKey(accessKey)
	if err != nil {
		l.Logger.Errorf("用户读取失败，%s", err.Error())
		return nil, fmt.Errorf("用户读取失败, %s", err.Error())
	}

	if err = r.ParseMultipartForm(32 << 20); err != nil {
		l.Logger.Errorf("上传时解析表单错误，%s", err.Error())
		return nil, fmt.Errorf("上传时解析表单错误, %s", err.Error())
	}

	file, fileHeader, err := r.FormFile("filename")
	if err != nil {
		l.Logger.Errorf("加载文件失败，%s", err.Error())
		return nil, fmt.Errorf("加载文件失败, %s", err.Error())
	}

	l.Logger.Infof("Object.receive upload file:%s", fileHeader.Filename)

	ctx := context.Background()
	bucketName := types.BucketName(user.TenantId, user.Namespace)
	fileName := fmt.Sprintf("%s_%s%s", time.Now().In(l.svcCtx.Loc).Format("2006_01_02.15.04"),
		hash.StringHash(fileHeader.Filename, 16), path.Ext(fileHeader.Filename))
	objectName := fmt.Sprintf("%s/%s/%s", types.BucketHome(accessKey), utils.PathFilter(uploadPath),
		fileName)
	_, err = l.svcCtx.Client.PutObject(ctx, bucketName, objectName, file, fileHeader.Size, client.PutObjectOptions{})
	if err != nil {
		l.Logger.Errorf("写入文件失败，%s", err.Error())
		return nil, fmt.Errorf("写入文件失败, %s", err.Error())
	}

	return &types.UploadResp{
		Url: fmt.Sprintf("http://%s/object/download?bucket=%s&objectID=%s&path=%s",
			l.svcCtx.Config.StorageHost, bucketName, objectName, uploadPath),
	}, nil
}
