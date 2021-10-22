package logic

import (
	"context"
	"fmt"
	"github.com/lifezq/minio-s3/api/client"
	"github.com/lifezq/minio-s3/api/internal/svc"
	"github.com/lifezq/minio-s3/api/internal/types"
	"github.com/lifezq/minio-s3/api/model"
	"github.com/lifezq/minio-s3/api/utils"
	"github.com/rs/xid"
	"os"
	"os/exec"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
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

	l.Logger.Infof("Register.receive req:%+v\n", req)
	accessKey := xid.New().String()
	secretKey := utils.GenerateSecretKey(accessKey)
	ctx := context.Background()
	err := exec.CommandContext(ctx, "mc", []string{"admin", "user", "add",
		l.svcCtx.Config.Minio.ServerName, accessKey, secretKey}...).Run()
	if err != nil {
		l.Logger.Errorf("创建用户失败：%s\n", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("创建用户失败,%s", err.Error())
	}

	bucket := types.BucketName(req.TenantID, req.Namespace)
	exists, err := l.svcCtx.Client.BucketExists(ctx, bucket)
	if err != nil || !exists {
		err = l.svcCtx.Client.MakeBucket(ctx, bucket, client.MakeBucketOptions{})
		if err != nil {
			l.Logger.Errorf("创建bucket失败：%s\n", err.Error())
			return &types.RegisterResp{}, fmt.Errorf("创建bucket失败,%s", err.Error())
		}
	}

	policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["s3:*"],"Resource":["arn:aws:s3:::` + bucket + `/` + types.Home(accessKey) + `/*"]}]}`
	policyFile := fmt.Sprintf("./%s.policy", accessKey)
	fp, err := os.OpenFile(policyFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 644)
	if err != nil {
		l.Logger.Errorf("写用户访问策略失败：%s\n", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("写用户访问策略失败,%s", err.Error())
	}

	fp.WriteString(policy)
	fp.Close()
	defer os.Remove(policyFile)

	err = exec.CommandContext(ctx, "mc", []string{"admin", "policy", "add",
		l.svcCtx.Config.Minio.ServerName, types.PolicyName(req.TenantID, req.Namespace, req.UserID), policyFile}...).Run()
	if err != nil {
		l.Logger.Errorf("添加访问策略失败：%s\n", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("添加访问策略失败,%s", err.Error())
	}

	err = exec.CommandContext(ctx, "mc", []string{"admin", "group", "add",
		l.svcCtx.Config.Minio.ServerName, types.GroupName(req.TenantID, req.Namespace), accessKey}...).Run()
	if err != nil {
		l.Logger.Errorf("用户组添加用户失败：%s\n", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("用户组添加用户失败,%s", err.Error())
	}

	err = exec.CommandContext(ctx, "mc", []string{"admin", "policy", "set",
		l.svcCtx.Config.Minio.ServerName, types.PolicyName(req.TenantID, req.Namespace, req.UserID), "user=" + accessKey}...).Run()
	if err != nil {
		l.Logger.Errorf("设置用户权限失败：%s\n", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("设置用户权限失败,%s", err.Error())
	}

	id, err := l.svcCtx.UserModel.FindMaxId()
	if err != nil {
		l.Logger.Errorf("获取用户最大id失败, %s\n", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("获取用户最大id失败,%s", err.Error())
	}

	_, err = l.svcCtx.UserModel.Insert(model.User{
		Id:        id + 1,
		AccessKey: accessKey,
		SecretKey: secretKey,
		TenantId:  req.TenantID,
		Namespace: req.Namespace,
		UserId:    req.UserID,
		CreateAt:  time.Now().In(l.svcCtx.Loc),
	})
	if err != nil {
		l.Logger.Errorf("创建用户失败：%s\n", err.Error())
		return &types.RegisterResp{}, fmt.Errorf("创建用户失败,%s", err.Error())
	}

	l.Logger.Infof("创建用户[%s]成功\n", accessKey)

	return &types.RegisterResp{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}, nil
}
