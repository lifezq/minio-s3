package logic

import (
	"context"
	"fmt"
	"time"

	"gitlab.energy-envision.com/storage/client"
	"gitlab.energy-envision.com/storage/internal/svc"
	"gitlab.energy-envision.com/storage/internal/types"
	"gitlab.energy-envision.com/storage/model"
	"gitlab.energy-envision.com/storage/utils"

	"github.com/rs/xid"
	"github.com/tal-tech/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (*types.RegisterResp, error) {

	l.Logger.Infof("Register.receive req:%+v", req)

	if !types.ValidNamespace(l.svcCtx.Config.ValidNamespace, req.Namespace) {
		l.Logger.Errorf("未注册的Namespace[%s]，请联系管理员, %s", req.Namespace)
		return &types.RegisterResp{}, fmt.Errorf("未注册的Namespace[%s]，请联系管理员, %s", req.Namespace)
	}

	accessKey := xid.New().String()
	secretKey := utils.GenerateSecretKey(accessKey)
	hashSecretKey, err := bcrypt.GenerateFromPassword([]byte(secretKey), 10)
	if err != nil {
		l.Logger.Errorf("用户secretKey生成错误, %s", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("用户secretKey生成错误,%s", err.Error())
	}

	ctx := context.Background()
	err = l.svcCtx.Cmd.CreateUser(ctx, l.svcCtx.Config.Minio.ServerName, accessKey, secretKey)
	if err != nil {
		l.Logger.Errorf("创建用户失败：%s", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("创建用户失败,%s", err.Error())
	}

	bucket := types.BucketName(req.TenantID, req.Namespace)
	exists, err := l.svcCtx.Client.BucketExists(ctx, bucket)
	if err != nil || !exists {
		err = l.svcCtx.Client.MakeBucket(ctx, bucket, client.MakeBucketOptions{})
		if err != nil {
			l.Logger.Errorf("创建bucket失败：%s", err.Error())
			return &types.RegisterResp{}, fmt.Errorf("创建bucket失败,%s", err.Error())
		}
	}

	err = l.svcCtx.Cmd.AddPolicy(ctx, l.svcCtx.Config.Minio.ServerName, bucket, types.BucketPolicyName(req.TenantID, req.Namespace, req.UserID), accessKey)
	if err != nil {
		l.Logger.Errorf("添加访问策略失败：%s", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("添加访问策略失败,%s", err.Error())
	}

	err = l.svcCtx.Cmd.AddGroupUser(ctx, l.svcCtx.Config.Minio.ServerName, types.BucketGroupName(req.TenantID, req.Namespace), accessKey)
	if err != nil {
		l.Logger.Errorf("用户组添加用户失败：%s", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("用户组添加用户失败,%s", err.Error())
	}

	id, err := l.svcCtx.UserModel.FindMaxId()
	if err != nil {
		l.Logger.Errorf("获取用户最大id失败, %s", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("获取用户最大id失败,%s", err.Error())
	}

	_, err = l.svcCtx.UserModel.Insert(model.User{
		Id:        id + 1,
		AccessKey: accessKey,
		SecretKey: string(hashSecretKey),
		TenantId:  req.TenantID,
		Namespace: req.Namespace,
		UserId:    req.UserID,
		CreateAt:  time.Now().In(l.svcCtx.Loc),
	})
	if err != nil {
		l.Logger.Errorf("创建用户失败：%s", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("创建用户失败,%s", err.Error())
	}

	l.Logger.Infof("创建用户[%s]成功", accessKey)

	return &types.RegisterResp{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}, nil
}
