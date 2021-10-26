package cmd

import "context"

type LocalCommand struct {
}

func NewLocalCommand() *LocalCommand {
	return &LocalCommand{}
}

func (m *LocalCommand) CreateUser(ctx context.Context, target, accessKey, secretKey string) error {
	//TODO
	return nil
}

func (m *LocalCommand) AddPolicy(ctx context.Context, target, bucket, policyName, accessKey string) error {
	//TODO
	return nil
}

func (m *LocalCommand) AddGroupUser(ctx context.Context, target, groupName, accessKey string) error {
	//TODO
	return nil
}
