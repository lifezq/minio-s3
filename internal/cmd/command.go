package cmd

import "context"

type Command interface {
	CreateUser(ctx context.Context, target, accessKey, secretKey string) error
	AddPolicy(ctx context.Context, target, bucket, policyName, accessKey string) error
	AddGroupUser(ctx context.Context, target, groupName, accessKey string) error
}
