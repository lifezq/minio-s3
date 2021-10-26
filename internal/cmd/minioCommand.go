package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"gitlab.energy-envision.com/storage/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MinioCommand struct {
}

func NewMinioCommand() *MinioCommand {
	return &MinioCommand{}
}

func (m *MinioCommand) CreateUser(ctx context.Context, target, accessKey, secretKey string) error {

	err := exec.CommandContext(ctx, "mc", []string{"admin", "user", "add",
		target, accessKey, secretKey}...).Run()
	if err != nil {
		logx.Errorf("创建用户失败：%s", err.Error())
		return fmt.Errorf("创建用户失败,%s", err.Error())
	}

	return nil
}

func (m *MinioCommand) AddPolicy(ctx context.Context, target, bucket, policyName, accessKey string) error {

	policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["s3:*"],"Resource":["arn:aws:s3:::` + bucket + `/` + types.BucketHome(accessKey) + `/*"]}]}`
	policyFile := fmt.Sprintf("./%s.policy", accessKey)
	fp, err := os.OpenFile(policyFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 644)
	if err != nil {
		logx.Errorf("写用户访问策略失败：%s", err.Error())
		return fmt.Errorf("写用户访问策略失败,%s", err.Error())
	}

	fp.WriteString(policy)
	fp.Close()
	defer os.Remove(policyFile)

	err = exec.CommandContext(ctx, "mc", []string{"admin", "policy", "add",
		target, policyName, policyFile}...).Run()
	if err != nil {
		logx.Errorf("添加访问策略失败：%s", err.Error())
		return fmt.Errorf("添加访问策略失败,%s", err.Error())
	}

	err = exec.CommandContext(ctx, "mc", []string{"admin", "policy", "set",
		target, policyName, "user=" + accessKey}...).Run()
	if err != nil {
		logx.Errorf("设置用户权限失败：%s", err.Error())
		return fmt.Errorf("设置用户权限失败,%s", err.Error())
	}

	return nil
}

func (m *MinioCommand) AddGroupUser(ctx context.Context, target, groupName, accessKey string) error {
	err := exec.CommandContext(ctx, "mc", []string{"admin", "group", "add",
		target, groupName, accessKey}...).Run()
	if err != nil {
		logx.Errorf("用户组添加用户失败：%s", err.Error())
		return fmt.Errorf("用户组添加用户失败,%s", err.Error())
	}
	return nil
}
