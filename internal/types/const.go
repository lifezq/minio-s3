package types

import (
	"fmt"
)

const (
	BUCKET_PREFIX        = "buk"
	BUCKET_GROUP_PREFIX  = "group"
	BUCKET_POLICY_PREFIX = "policy"
	BUCKET_HOME_PREFIX   = "home"

	S3_AUTHORIZATION = "S3Authorization"

	CACHE_REDIS_STATE = "crs"

	ENGINE_LOCAL = "Local"
	ENGINE_MINIO = "MiniO"

	META_VERSION = "v0"
)

func BucketName(tenantID, nameSpace string) string {
	return fmt.Sprintf("%s.%s.%s", BUCKET_PREFIX, tenantID, nameSpace)
}

func BucketGroupName(tenantID, nameSpace string) string {
	return fmt.Sprintf("%s.%s.%s", BUCKET_GROUP_PREFIX, tenantID, nameSpace)
}

func BucketPolicyName(tenantID, nameSpace, userID string) string {
	return fmt.Sprintf("%s.%s.%s.%s", BUCKET_POLICY_PREFIX, tenantID, nameSpace, userID)
}

func BucketHome(s string) string {
	return fmt.Sprintf("%s.%s", BUCKET_HOME_PREFIX, s)
}

func BucketUserObjectPath(accessKey string, path, object string) string {
	return fmt.Sprintf("%s.%s/%s/%s", BUCKET_HOME_PREFIX, accessKey, path, object)
}
