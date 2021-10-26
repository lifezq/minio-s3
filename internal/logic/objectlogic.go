package logic

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"gitlab.energy-envision.com/storage/client"
	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"
	"gitlab.energy-envision.com/storage/utils"

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
	uploadPath := utils.PathFilter(token["path"].(string))

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
	objectID := hash.StringHash(fmt.Sprintf("%d.%s", fileHeader.Size, fileHeader.Filename), 16)
	objectName := fmt.Sprintf("%s/%s/%s", types.BucketHome(accessKey), uploadPath, objectID)

	_, err = l.svcCtx.Client.StatObject(ctx, bucketName, objectName, client.StatObjectOptions{})
	if err != nil {
		l.setObjectPathSizeAndCount(ctx, bucketName, fmt.Sprintf("%s/%s/",
			types.BucketHome(accessKey), uploadPath), fileHeader.Size)

		l.setObjectCount(ctx, bucketName, fmt.Sprintf("%s/", types.BucketHome(accessKey)))
	}

	_, err = l.svcCtx.Client.PutObject(ctx, bucketName, objectName, file, fileHeader.Size, client.PutObjectOptions{
		UserMetadata: map[string]string{
			"filename":     fileHeader.Filename,
			"meta-version": types.META_VERSION},
	})
	if err != nil {
		l.Logger.Errorf("写入文件失败，%s", err.Error())
		return nil, fmt.Errorf("写入文件失败, %s", err.Error())
	}

	return &types.UploadResp{
		Url: fmt.Sprintf("http://%s/object/download?bucket=%s&objectID=%s&path=%s",
			l.svcCtx.Config.StorageHost, bucketName, objectID, uploadPath),
	}, nil
}

func (l *ObjectLogic) setObjectPathSizeAndCount(ctx context.Context, bucketName, objectName string, size int64) {

	pathInfo, _ := l.svcCtx.Client.StatObject(ctx, bucketName, objectName, client.StatObjectOptions{})
	folderSize, ok := pathInfo.UserMetadata["Folder-Size"]
	if ok {
		sz, _ := strconv.ParseUint(folderSize, 10, 64)
		folderSize = strconv.FormatUint(sz+uint64(size), 10)
	} else {
		folderSize = strconv.FormatInt(size, 10)
	}

	objectCount, ok := pathInfo.UserMetadata["Object-Count"]
	if ok {
		count, _ := strconv.ParseUint(objectCount, 10, 64)
		objectCount = strconv.FormatUint(count+1, 10)
	} else {
		objectCount = strconv.FormatInt(1, 10)
	}

	l.svcCtx.Client.PutObject(ctx, bucketName, objectName, strings.NewReader(""), 0, client.PutObjectOptions{
		UserMetadata: map[string]string{
			"folder-size":  folderSize,
			"object-count": objectCount,
			"meta-version": types.META_VERSION,
		}})
}

func (l *ObjectLogic) setObjectCount(ctx context.Context, bucketName, objectName string) {
	// path size
	pathInfo, _ := l.svcCtx.Client.StatObject(ctx, bucketName, objectName, client.StatObjectOptions{})
	objectCount, ok := pathInfo.UserMetadata["Object-Count"]
	if ok {
		count, _ := strconv.ParseUint(objectCount, 10, 64)
		objectCount = strconv.FormatUint(count+1, 10)
	} else {
		objectCount = strconv.FormatInt(1, 10)
	}

	l.svcCtx.Client.PutObject(ctx, bucketName, objectName, strings.NewReader(""), 0, client.PutObjectOptions{
		UserMetadata: map[string]string{
			"object-count": objectCount,
			"meta-version": types.META_VERSION,
		}})
}
