package logic

import (
	"context"
	"fmt"
	"strconv"

	"gitlab.energy-envision.com/storage/client"
	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"
	"gitlab.energy-envision.com/storage/utils"

	"github.com/tal-tech/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type TokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokenLogic {
	return TokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenLogic) Token(req types.TokenReq) (*types.TokenResp, error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("Token.receive req:%+v", req)

	user, err := l.svcCtx.UserModel.FindOneByAccessKey(req.AccessKey)
	if err != nil {
		l.Logger.Errorf("获取用户失败：%s\n", err.Error())
		return &types.TokenResp{}, fmt.Errorf("获取用户失败,%s", err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.SecretKey), []byte(req.SecretKey))
	if err != nil {
		l.Logger.Errorf("用户secretKey验证不通过：%s", err.Error())
		return &types.TokenResp{}, fmt.Errorf("用户secretKey验证不通过")
	}

	folderInfo, err := l.svcCtx.Client.StatObject(context.Background(), types.BucketName(user.TenantId, user.Namespace),
		fmt.Sprintf("%s/%s/", types.BucketHome(req.AccessKey),
			utils.PathFilter(req.Path)), client.StatObjectOptions{})
	if objectCount, ok := folderInfo.UserMetadata["Object-Count"]; ok {
		count, _ := strconv.Atoi(objectCount)
		if uint16(count) > l.svcCtx.Config.FolderMaxFiles {
			l.Logger.Errorf("该路径下文件超过限制")
			return &types.TokenResp{}, fmt.Errorf("该路径下文件超过限制")
		}
	}

	s3Authorization, err := types.CreateToken(req.AccessKey, req.SecretKey, req.Path, l.svcCtx.Config.TokenExpireIn)
	if err != nil {
		l.Logger.Errorf("创建令牌发生错误：%s", err.Error())
		return &types.TokenResp{}, fmt.Errorf("创建令牌发生错误")
	}

	return &types.TokenResp{
		S3Authorization: s3Authorization,
		ExpireIn:        l.svcCtx.Config.TokenExpireIn,
	}, nil
}
